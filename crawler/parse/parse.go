package parse

import (
	"crawler/engine"
	"regexp"
)

func ParseTag(content []byte) engine.ParseResult {
	reg := regexp.MustCompile(`<a href="([^"]+)" class="tag">([^<""]+)</a>`)
	match := reg.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}

	for _, m := range match {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:       "https://book.douban.com" + string(m[1]),
			ParseFunc: ParseBookList,
		})
	}

	return result
}
