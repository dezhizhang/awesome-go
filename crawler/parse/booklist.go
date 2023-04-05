package parse

import (
	"crawler/engine"
	"fmt"
	"regexp"
)

func ParseBookList(contents []byte) engine.ParseResult {

	reg := regexp.MustCompile(`<a href="([^"]+)" title="([^"]+)"`)

	match := reg.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	fmt.Println(result)

	for _, m := range match {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: engine.NilParse,
		})
	}
	return result
}
