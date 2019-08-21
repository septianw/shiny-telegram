package main

import (
	"database/sql"
	// "errors"

	"fmt"

	_ "gopkg.in/goracle.v2"
)

// func New() Database {
// 	return Database{
// 		PingDb: func(d Dbconf) (bool, error) {
// 			err := errors.New("fail")
// 			return true, err
// 		},
// 		SetupDb: func(d Dbconf) bool {
// 			return true
// 		},
// 		Migrate: func(s string, d Dbconf) {
// 			return
// 		},
// 		OpenDb: func(d Dbconf) (*sql.DB, error) {
// 			return
// 		},
// 	}
// }

func PingDb(d Dbconf) (bool, error) {

	dsn := fmt.Sprintf("%s/%s@%s:%d/%s", d.User, d.Pass, d.Host, d.Port, d.Database)

	dbi, err := sql.Open("goracle", dsn)
	if err != nil {
		return false, err
	}

	err = dbi.Ping()
	if err != nil {
		return false, err
	}

	return true, nil
}

func OpenDb(d Dbconf) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s/%s@%s:%d/%s",
		d.User, d.Pass, d.Host, d.Port, d.Database)

	database, err := sql.Open("goracle", dsn)

	return database, err
}

func Migrate(location string, d Dbconf) {
	return
}

func SetupDb(d Dbconf) bool {
	return true
}
