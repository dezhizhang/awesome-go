package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "abc a7c mfc cat cba"

	ret := regexp.MustCompile("a.c")

	all := ret.FindAllStringSubmatch(str, -1)
	fmt.Println("all", all)

}
