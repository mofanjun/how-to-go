package main

import (
	"crawler/scheduler/engine"
	"crawler/scheduler/scheduler"
	"crawler/scheduler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount:10,
	}


	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
