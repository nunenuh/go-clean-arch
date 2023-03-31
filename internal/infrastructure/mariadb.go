package infrastructure

import (
	"database/sql"
	"time"
)

func ConnectToMariaDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(mysql:3306)/fiber_post")
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
