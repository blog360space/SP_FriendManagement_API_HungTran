package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"flag"
)

func DbConn() (db *sql.DB) {
	var dbName string = "friend_managerment"
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

func TruncateTable(table string) (bool, error) {
	db := DbConn();
	_, err := db.Query("TRUNCATE TABLE " + table)

	return true, err
}