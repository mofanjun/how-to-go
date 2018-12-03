package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	WorkerReady(chan Request)
	Run()
	ConfigMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		creteWorker(out,e.Scheduler)
	}

	for _,r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <- out
		//deal item
		for _,item := range result.Items {
			log.Printf("Got item #%d: %v",itemCount,item)
			itemCount++
		}

		//deal request
		for _,request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func creteWorker(out chan ParseResult,s Scheduler)  {
	go func() {
		in := make(chan Request)
		for  {
			//告诉调度器空闲了
			s.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}