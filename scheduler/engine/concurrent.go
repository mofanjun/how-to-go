package engine

import "log"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		creteWorker(in,out)
	}

	for _,r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <- out
		//deal item
		for _,item := range result.Items {
			log.Printf("Got item %v",item)
		}

		//deal request
		for _,request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func creteWorker(in chan Request,out chan ParseResult)  {
	go func() {
		for  {
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}