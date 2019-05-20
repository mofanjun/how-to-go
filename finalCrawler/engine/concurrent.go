package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemSaver chan Item
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
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <- out
		//deal item
		for _,item := range result.Items {
			go func() {e.ItemSaver <- item}()
		}

		//deal request
		for _,request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
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



var visitedUrl = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}

	visitedUrl[url] = true
	return false
}
