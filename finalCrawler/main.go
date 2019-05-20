package main

import (
	"crawler/finalCrawler/engine"
	"crawler/finalCrawler/scheduler"
	"crawler/finalCrawler/zhanai/parser"
	"crawler/finalCrawler/persist"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")

	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount:10,
		ItemSaver:itemChan,
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
