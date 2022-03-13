package repository

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8","root", "falcon2002", "localhost", "3306", "test")
    Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
        fmt.Println("falcon")
		panic(err)
	}
}

