package main

import (
	"crawler/finalCrawler/engine"
	"crawler/finalCrawler/scheduler"
	"crawler/finalCrawler/zhanai/parser"
	"crawler/finalCrawler/persist"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount:10,
		ItemSaver:persist.ItemSaver(),
	}


	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun/hangzhou",
		ParseFunc:parser.ParseCity,
	})
}
