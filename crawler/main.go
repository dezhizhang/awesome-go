package main

import (
	"crawler/engine"
	"crawler/parse"
)

func main() {
	//engine.Run(engine.Request{
	//	Url:       "https://book.douban.com/",
	//	ParseFunc: parse.ParseTag,
	//})

	concurrent := engine.ConcurrentEngine{
		Scheduler: &engine.SimpleScheduler{},
		WorkCount: 100,
	}

	concurrent.Run(engine.Request{
		Url:       "https://book.douban.com/",
		ParseFunc: parse.ParseTag,
	})

}
