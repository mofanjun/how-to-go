package fetcher

import (
	"time"
	"net/http"
	"fmt"
	"bufio"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
)

//rateLimiter 在包下面是全局的 所有的fetcher共享
var rateLimiter = time.Tick(20 * time.Millisecond)

//
func Fetch(url string) ([]byte , error) {
	<- rateLimiter
	resp, err := agentGet(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d",resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding (r *bufio.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}

func agentGet (url string) (*http.Response, error){
	client := &http.Client{}
	request, err := http.NewRequest("GET",url,nil)

	if err != nil {
		return nil, err
	}
	//add user agent
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.157 Safari/537.36")
	resp,err := client.Do(request)

	return resp,nil
}
