package dboracle

import (
	"database/sql"
	// "errors"

	"fmt"

	pak "github.com/septianw/shiny-telegram/experiment01/sharedpak"
	_ "gopkg.in/goracle.v2"
)

func PingDb(d pak.Dbconf) (bool, error) {

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

func OpenDb(d pak.Dbconf) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s/%s@%s:%d/%s",
		d.User, d.Pass, d.Host, d.Port, d.Database)

	database, err := sql.Open("goracle", dsn)

	return database, err
}

func Migrate(location string, d pak.Dbconf) {
	return
}

func SetupDb(d pak.Dbconf) bool {
	return true
}
