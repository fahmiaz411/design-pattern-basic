package singleton

import (
	"database/sql"
	"log"
)

type OrderService struct {
	DB *sql.DB
}

func (s *OrderService) Save(id string){
	_, err := s.DB.Exec("INSERT INTO orders ...")
	if err != nil {
		log.Println(err)
	}
}