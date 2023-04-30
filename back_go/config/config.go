package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

const (
	DbDriver   = "mysql"
	DbUserName = "root"
	DbPasswd   = "12345"
	DbHost     = "127.0.0.1"
	DbPort     = 3306
	DbName     = "test"
)

func InitDB() {
	var err error
	Db, err = sql.Open(DbDriver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		DbUserName, DbPasswd, DbHost, DbPort, DbName))
	if err != nil {
		panic(err)
	}
	Db.SetConnMaxLifetime(100 * time.Second)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(16)
}
