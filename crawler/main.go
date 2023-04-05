package main

import (
	"crawler/engine"
	"crawler/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:       "https://book.douban.com/",
		ParseFunc: parse.ParseTag,
	})
}
