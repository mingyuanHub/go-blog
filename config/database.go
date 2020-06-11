package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Db() *sql.DB{
	db, _ = sql.Open("mysql", "root:123456@tcp(39.100.133.182:3306)/goBlog?charset=utf8")
	return db
}

