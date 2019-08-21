package main

import (
	"database/sql"
)

type PingDbFunc func(d Dbconf) (bool, error)
type SetupDbFunc func(d Dbconf) bool
type MigrateFunc func(location string, d Dbconf)
type OpenDbFunc func(d Dbconf) (*sql.DB, error)

type databaseType string

type DatabaseInterface interface {
	PingDb(Dbconf) (bool, error)
	SetupDb(Dbconf) bool
	Migrate(string, Dbconf)
	OpenDb(Dbconf) (*sql.DB, error)
	// PingDb  PingDbFunc
	// SetupDb SetupDbFunc
	// Migrate MigrateFunc
	// OpenDb  OpenDbFunc
	// Db      *sql.DB
}
