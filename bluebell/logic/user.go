package logic

import (
	"bluebell/dao"
	"fmt"
	"log"
)

func CheckUserExit(email string) (userId int64, err error) {
	var id int64
	sqlStr := "select * from user where email = ?"
	fmt.Println(email)
	row := dao.DB.QueryRow(sqlStr, &email)
	err = row.Scan(&id)
	if err != nil {
		log.Printf("扫描失败")
		return 0, err
	}
	return id, nil
}
