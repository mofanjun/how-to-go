package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"regexp"
)

const fetchUrl  = "http://www.zhenai.com/zhenghun"
//<a href="http://www.zhenai.com/zhenghun/anshun" class="">安顺</a>

func main() {
	resp, err := http.Get(fetchUrl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code",resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)

	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n",all)
	parseCityList(all)
}

func parseCityList (contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City: %s URL: %s\n",m[2],m[1])
	}

	fmt.Printf("Find City total count %d",len(matches))
}

func determineEncoding (r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
