package main

import (
"crawler/singleCrawler/engine"
"crawler/singleCrawler/zhenai/parser"
	engine2 "crawler/scheduler/engine"
)

func main() {
	engine2.SimpleEngine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
