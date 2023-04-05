package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {

	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://book.douban.com/", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")

	if err != nil {
		panic(err)
	}

	response, _ := client.Do(request)
	defer response.Body.Close()

	//fmt.Println(rsp.StatusCode)
	//
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d", response.StatusCode)
	}

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	parseContent(result)
}

func parseContent(content []byte) {
	reg := regexp.MustCompile(`<a href="([^"]+)" class="tag">([^""]+)</a>`)
	match := reg.FindAllSubmatch(content, -1)

	for _, m := range match {
		fmt.Printf("m[0]:%s m[1]:%s m[2]:%s", m[0], m[1], m[2])
	}
}
