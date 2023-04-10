package test

import (
	"bluebell/utils"
	"fmt"
	"log"
	"testing"
)

func TestNewNode(t *testing.T) {
	id, err := utils.NewNode()
	if err != nil {
		log.Printf("生成失败")
	}

	fmt.Println(id)
}
