package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"flag"
)

func DbConn() (db *sql.DB) {
	var dbName string = "bar_tracking"
	var dbDriver string = "mysql"
	var dbUser string = "root"
	var dbPass string = "123456"

	if flag.Lookup("test.v") != nil {
		dbName += "_test"
	}

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}
