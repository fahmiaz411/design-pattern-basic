package singleton

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


func connect(host, username, password string) *sql.DB {	
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", username, password, host, 3306, "test"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}