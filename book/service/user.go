package service

import "book/dao"

func AddUser() error {
	sqlStr := "insert into users(id,name,password,email) values (?,?,?,?)"
	result, err := dao.DB.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}

	_, err = result.Exec("456", "张三", "13456", "1541609448@qq.com")
	if err != nil {
		panic(err)
	}
	return nil
}

//func AddUser() error {
//	sqlStr := "insert into users(id,name,password,email) values (?,?,?,?)"
//
//	result, err := dao.DB.Prepare(sqlStr)
//	if err != nil {
//		log.Panic(err)
//	}
//
//	_, err2 := result.Exec("123", "张三", "123456", "1541@qq.com")
//	if err2 != nil {
//		panic(err)
//	}
//	return nil
//}
