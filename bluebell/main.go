package main

import "bluebell/router"

func main() {
	r := router.StepRouter()

	r.Run(":8080")
}
