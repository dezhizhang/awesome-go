package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	DB  *sql.DB
	err error
)

func init() {
	DB, err = sql.Open("mysql", "root:701XTAY1993@/bluebell")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
