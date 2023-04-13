package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	dsn := "root:701XTAY1993@tcp(127.0.0.1:3306)/crm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("边接数据库失败%s", err)
		return
	}
	DB = db
}
