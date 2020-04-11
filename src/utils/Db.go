package utils

import (
	"flag"

	// No use
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var dbConn *gorm.DB

// DbConn Connect to db
func DbConn() *gorm.DB {
	if dbConn != nil {
		return dbConn
	}

	var dbName string = "friend_managerment"
	var dbDriver string = "mysql"
	var dbUser string = "root"
	var dbPass string = "123456"

	if flag.Lookup("test.v") != nil {
		dbName += "_test"
	}

	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	dbConn, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}
	return dbConn
}

// DbTruncateTable Truncate specify table
func DbTruncateTable(table string) bool {
	db := DbConn()

	db.Exec("TRUNCATE TABLE " + table)

	return true
}

// DbClose Close Db Connection
func DbClose() {
	db := DbConn()
	db.Close()
}
