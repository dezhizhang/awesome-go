package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "1541609448@qq.com"

	reg := regexp.MustCompile(str)

	res := reg.FindString(str)
	fmt.Println("all", res)

}
