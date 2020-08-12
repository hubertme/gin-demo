package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDriver() {
	fmt.Println("MySQL init driver running")

	dbTmp, err := sql.Open("mysql", "root:testtist@tcp(127.0.0.1:3306)/test_db")

	if err != nil {
		panic(err.Error())
	}

	DB = dbTmp

	//defer DB.Close()
}
