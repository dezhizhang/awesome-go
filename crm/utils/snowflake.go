package utils

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

func SnowflakeId() (int64, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Printf("生成雪花算法失败")
		return 0, err
	}

	id := node.Generate().Int64()
	return id, nil
}
