package engine

import "log"
import "crawler/finalCrawler/fetcher"

func worker(r Request) (ParseResult, error){
	//log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch error fetching url %s %v",r.Url,err)
		return ParseResult{}, nil
	}

	return r.ParseFunc(body,r.Url), nil
}
