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

	//engine.Run(engine.Request{
	//	Url:       "https://book.douban.com/tag/%E9%9A%8F%E7%AC%94",
	//	ParseFunc: parse.ParseBookList,
	//})

	engine.Run(engine.Request{
		Url:       "https://book.douban.com/subject/36177518/",
		ParseFunc: parse.ParseBookDetail,
	})
}
