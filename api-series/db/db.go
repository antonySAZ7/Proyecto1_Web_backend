package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := "user=postgres dbname=mi_seriesdb_24710 sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error conectando a DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB no responde:", err)
	}

	log.Println("conectado a pg")
}
