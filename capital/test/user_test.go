package test

import (
	"book/service"
	"testing"
)

func TestAddUser(t *testing.T) {

	service.SaveUser("刘德华", "123456", "1541609448@qq.com")
}

//func TestGetUser(t *testing.T) {
//	query, err := service.GetUserQuery()
//	if err != nil {
//		log.Printf("获取结果失败")
//	}
//	fmt.Println(query)
//}

//func TestGetUsers(t *testing.T) {
//	query, err := service.GetUsersQuery()
//	if err != nil {
//		log.Printf("获取结果失败")
//	}
//	fmt.Println(query)
//}
