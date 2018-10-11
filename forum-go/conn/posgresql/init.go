package posgresql

import (
	"database/sql"
	"fmt"
)

type DB struct {
	Name     string
	Addr     string
	User     string
	Password string
}

func InitDB(dbin DB) (dbout *sql.DB, err error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbin.User, dbin.Password, dbin.Name)
	db, err := sql.Open("postgres", dbinfo)
	return db, err
}
