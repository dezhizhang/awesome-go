package parse

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
)

var authorReg = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var pressReg = regexp.MustCompile(`<span class="pl"> 出版社：</span>([^<]+)<br>`)
var pageReg = regexp.MustCompile(`<span class="pl">页数:</span>([^<]+)<br>`)
var priceReg = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br>`)
var scoreReg = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)

func ParseBookDetail(contents []byte) engine.ParseResult {

	detail := model.BookDetail{}
	detail.Author = ExtraString(contents, authorReg)
	detail.Press = ExtraString(contents, pressReg)
	detail.BookPage = ExtraString(contents, pageReg)
	detail.Price = ExtraString(contents, priceReg)
	detail.Score = ExtraString(contents, scoreReg)

	result := engine.ParseResult{
		Items: []interface{}{detail},
	}
	return result

}

func ExtraString(contents []byte, reg *regexp.Regexp) string {
	match := reg.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
