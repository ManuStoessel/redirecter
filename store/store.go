package store

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	DB *sql.DB
}

func InitializeStore() *Store {
	db, err := sql.Open("sqlite3", "redirecter.sqlite")
	if err != nil {
		log.Panic(err)
	}

	tblStmt := "CREATE TABLE IF NOT EXISTS data (shorturl TEXT not null primary key, longurl TEXT, isodatetime TEXT);"
	_, err = db.Exec(tblStmt)
	if err != nil {
		log.Panic(err)
	}

	return &Store{DB: db}
}

func (s *Store) GetLongURL(shorturl string) (string, error) {
	sqlStmt := "SELECT longurl FROM data WHERE shorturl = ?"

	rows, err := s.DB.Query(sqlStmt, shorturl)
	if err != nil {
		return "", err
	}

	var longurl string

	if rows.Next() {
		err = rows.Scan(&longurl)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("No result")
	}

	return longurl, nil
}
