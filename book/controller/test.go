package controller

import (
	"book/dao"
	"book/model"
	"html/template"
	"log"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("./view/test.html"))
	sqlStr := "select * from users"
	rows, err := dao.DB.Query(sqlStr)
	if err != nil {
		log.Printf("查询失败%s", err)
	}
	var user model.User
	var users []model.User

	for rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.Gender,
			&user.Phone,
			&user.Address,
		)

		if err != nil {
			log.Printf("扫描失败%s", err)
		}
		users = append(users, user)
	}
	must.Execute(w, users)
}
