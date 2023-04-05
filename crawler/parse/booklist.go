package parse

import (
	"crawler/engine"
	"regexp"
)

func ParseBookList(contents []byte) engine.ParseResult {

	reg := regexp.MustCompile(`<a href="([^"]+)" title="([^"]+)"`)

	match := reg.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range match {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseBookDetail,
		})
	}
	return result
}
