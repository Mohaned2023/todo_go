package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func InitDB(url string) *sqlx.DB {
	DBConn, err := sqlx.Connect("postgres", url);
	if err != nil {
		panic(fmt.Errorf("%v", err))
	}

	DBConn.SetMaxOpenConns(20);
	DBConn.SetMaxIdleConns(5);
	return DBConn
}
