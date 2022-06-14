package objectpool

import (
	"database/sql"
	"design-pattern-basic/singleton"
	"errors"
)

var pool []*sql.DB

func getConnection() (*sql.DB, error) {
	if len(pool) == 0 {
		return nil, errors.New("End of connection")
	} 

	pool = pool[1:]
	return pool[0], nil
}

func close(db *sql.DB) {
	pool = append(pool, db)
}

func init() {
	for i := 0; i < 100; i++ {
		pool = append(pool, singleton.Connect("localhost", "root", "1234"))
	}
}