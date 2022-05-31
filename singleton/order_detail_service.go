package singleton

import (
	"database/sql"
	"log"
)

type OrderDetailService struct {
	DB *sql.DB
}

func (s *OrderDetailService) Save(id, product string) {
	_, err := s.DB.Exec("INSERT INTO orders ...")
	if err != nil {
		log.Println(err)
	}
}