package main

import (
	"crm/router"
	"crm/utils"
	"log"
)

func main() {

	// 初始化转换器
	err := utils.InitTrans("zh")
	if err != nil {
		log.Printf("转换失败")
	}

	r := router.StepRouter()

	r.Run(":8080")
}
