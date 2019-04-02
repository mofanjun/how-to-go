package engine

import "log"
import "crawler/finalCrawler/fetcher"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		creteWorker(e.Scheduler.WorkChan(),out,e.Scheduler)
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

func creteWorker(in chan Request,
	out chan ParseResult,r ReadyNotifier)  {
	go func() {
		for  {
			//告诉调度器空闲了
			r.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r Request) (ParseResult, error){
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch error fetching url %s %v",r.Url,err)
		return ParseResult{}, nil
	}

	return r.ParseFunc(body), nil
}
