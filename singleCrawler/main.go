package main

import (
"crawler/singleCrawler/engine"
"crawler/singleCrawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
