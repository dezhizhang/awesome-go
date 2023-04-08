package test

import (
	"book/service"
	"fmt"
	"log"
	"testing"
)

func TestAddUser(t *testing.T) {

	service.AddUser()
}

func TestGetUser(t *testing.T) {
	query, err := service.GetUserQuery()
	if err != nil {
		log.Printf("获取结果失败")
	}
	fmt.Println(query)
}

func TestGetUsers(t *testing.T) {
	query, err := service.GetUsersQuery()
	if err != nil {
		log.Printf("获取结果失败")
	}
	fmt.Println(query)
}
