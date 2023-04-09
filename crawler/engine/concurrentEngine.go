package engine

import (
	"crawler/fetch"
	"fmt"
	"log"
)

type Scheduler interface {
	Submit(request Request)
	Run()
	WorkReady(chan Request)
	WorkChan() chan Request
}

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type SimpleScheduler struct {
	WorkerChan chan Request
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	c.Scheduler.Run()

	for i := 0; i < c.WorkCount; i++ {
		CreateWork(c.Scheduler.WorkChan(), out, c.Scheduler)
	}

	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item:%s", item)
		}
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

func CreateWork(in chan Request, out chan ParseResult, s Scheduler) {
	go func() {
		for {
			s.WorkReady(in)
			request := <-in

			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func Worker(r Request) (ParseResult, error) {
	fmt.Printf("Fetcb url:%s", r.Url)

	body, err := fetch.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch Error: s%", r.Url)
		return ParseResult{}, err
	}

	return r.ParseFunc(body), nil
}
