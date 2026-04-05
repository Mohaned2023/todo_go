package storage

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DBConn *sqlx.DB;

func InitDB(url string) {
	var err error;
	DBConn, err = sqlx.Connect("postgres", url);
	if err != nil {
		log.Fatalf("Can not connect to database! %v", err)
	}

	DBConn.SetMaxOpenConns(20);
	DBConn.SetMaxIdleConns(5);
}
