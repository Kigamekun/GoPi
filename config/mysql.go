package config

import (
	"database/sql"
	"log"
	_"github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "django:ReksaSyahputra1012!@tcp(localhost:3306)/pegawai")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

