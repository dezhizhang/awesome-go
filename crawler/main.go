package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func httpGet(url string) (result string, err error) {
	rsp, err := http.Get(url)
	defer rsp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err2 := rsp.Body.Read(buf)
		if n == 0 {
			fmt.Println("数据读取失败")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			break
		}
		result += string(buf[:n])
	}
	return
}

func spiderPage(i int, page chan int) {
	url := "https://tieba.baidu.com/f?kw=golang&ie=utf-8&pn=" + strconv.Itoa(i+50) + "&pagelets=frs-list%2Fpagelet%2Fthread&pagelets_stamp=1680663932369"
	result, err := httpGet(url)
	if err != nil {
		fmt.Println("获取失败", err)
	}

	f, err := os.Create("第" + strconv.Itoa(i) + "页数据" + ".html")
	if err != nil {
		fmt.Println("写入文件失败")
		return
	}

	f.WriteString(result)
	f.Close()

	page <- i

}

func working(start int, end int) {
	fmt.Println("开始工作")

	page := make(chan int)

	for i := start; i <= end; i++ {

		go spiderPage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成", <-page)
	}

}

func main() {
	var start, end int
	fmt.Println("请输入攫取的起始页")
	fmt.Scan(&start)
	fmt.Println("请输入攫取的终止页")
	fmt.Scan(&end)

	working(start, end)
}
