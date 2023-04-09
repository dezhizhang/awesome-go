package utils

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

func NewNode() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Printf("生成失败")
	}

	// Generate a snowflake ID.
	id := node.Generate().Int64()
	return id
}
