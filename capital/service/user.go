package service

import (
	"book/dao"
	"book/model"
	"log"
)

func AddUser() error {
	sqlStr := "insert into users(id,name,password,email) values (?,?,?,?)"
	result, err := dao.DB.Prepare(sqlStr)
	if err != nil {
		log.Printf("预编译失败%s", err)
	}

	_, err = result.Exec("12312", "刘德华", "123456", "1541609448@qq.com")
	if err != nil {
		log.Printf("添加失败%s", err)
	}
	return nil
}

func GetUserQuery() (model.User, error) {
	sqlStr := "select * from users where id = ?"
	row := dao.DB.QueryRow(sqlStr, "456")

	user := model.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
	if err != nil {
		log.Printf("扫描失败%s", err)
	}
	return user, nil
}

func GetUsersQuery() ([]model.User, error) {
	sqlStr := "select * from users"
	rows, err := dao.DB.Query(sqlStr)
	if err != nil {
		log.Printf("查询失败%s", err)
	}
	var user model.User
	var users []model.User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		if err != nil {
			log.Printf("扫描失败%s", err)
		}
		users = append(users, user)
	}
	return users, nil
}
