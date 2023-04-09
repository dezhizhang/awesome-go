package test

import (
	"bluebell/utils"
	"fmt"
	"testing"
)

func TestNewNode(t *testing.T) {
	id := utils.NewNode()
	fmt.Println(id)
}
