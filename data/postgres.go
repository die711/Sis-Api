package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"io/ioutil"
	"os"
)

func getConnection() (*sql.DB, error) {
	uri := os.Getenv("DATABASE_URI")
	//uri := "postgres://postgres:11597@127.0.0.1:5432/sis?sslmode=disable"
	return sql.Open("postgres", uri)
}

func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("./database/models.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}