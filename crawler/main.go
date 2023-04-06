package main

import (
	"crawler/engine"
	"crawler/parse"
	"crawler/scheduler"
)

func main() {
	//engine.Run(engine.Request{
	//	Url:       "https://book.douban.com/",
	//	ParseFunc: parse.ParseTag,
	//})

	concurrent := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 100,
	}

	concurrent.Run(engine.Request{
		Url:       "https://book.douban.com/",
		ParseFunc: parse.ParseTag,
	})

}
