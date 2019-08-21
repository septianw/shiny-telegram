package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/gocraft/dbr"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

type Modules struct {
	Name      string
	Path      string
	Installed uint8
	Loaded    uint8
}

func PingDb(d Dbconf) (bool, error) {
	var reval bool = false
	var err error

	// try to connect
	Dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		d.User, d.Pass, d.Host, d.Port, d.Database)
	db, err := sql.Open(d.Type, Dsn)
	// fmt.Println(dsn)

	// if connect succeed, true, else false.
	if err != nil {
		// TODO: Handle error here
		reval = false
		fmt.Println(err)
	} else {
		reval = true
	}

	// connect succeed, lets try to ping it.
	err = db.Ping()
	if err != nil {
		// TODO: Handle error here.
		fmt.Println(err)
		reval = false
	} else {
		reval = true
	}

	// if all fail, do not fill Dsn.
	if !reval {
		Dsn = ""
	}

	return reval, err
}

func SetupDb(d Dbconf) bool {
	// TODO
	// Plan #1 Original plan
	// baca schema dari file system
	// terjemahkan schema ke create table
	// eksekusi create table.

	// Plan #2 Migration plan
	// baca migration file dari file system
	// migration up/down dari migration file.
	// buat dan baca schema dari struct
	var out bool = true
	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s?charset=utf8",
		d.User, d.Pass, d.Host, d.Port, d.Database)
	// m, err := migrate.New("file:///home/asep/gocode/src/github.com/septianw/shiny-telegram/experiment01/schema/main", dsn)
	m, err := migrate.New("file://./schema/main", dsn)
	// migrate.
	if err != nil {
		// TODO: Handle this error.
		log.Fatalln("wow")
		log.Fatal(err)
		out = false
	} else {
		out = true
	}

	// Migrate all the way up ...

	if err := m.Up(); err != nil {
		// TODO: Handle this error.
		// log.Fatalln("wow")
		if err == migrate.ErrNoChange {
			out = true
		} else {
			log.Fatal(err)
			out = false
		}
	} else {
		out = true
	}

	return out
}
