package singleton

import (
	"database/sql"
	"log"
)

type OrderServiceDetail struct {
	DB *sql.DB
}

func (s *OrderServiceDetail) Save() {
	_, err := s.DB.Exec("INSERT INTO orders ...")
	if err != nil {
		log.Println(err)
	}
}