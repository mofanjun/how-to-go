package parser

import (
	"regexp"
	"crawler/finalCrawler/engine"
	)


var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/hangzhou/[^"]+)"`)
)

func ParseCity (contents []byte) engine.ParseResult{
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		//clone it
		xurl := string(m[1])
		name := string(m[2])

		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c,name,xurl)
			},
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents,-1)
	for _, m := range matches {
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			ParseFunc:ParseCity,
		})
	}

	return result
}
