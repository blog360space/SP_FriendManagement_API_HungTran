package utils

import (
	"configs"
	"flag"
	"fmt"
	"os"
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

	var dbName string = configs.DB_NAME
	if flag.Lookup("test.v") != nil {
		dbName = configs.DB_NAME_TEST
	}

	dbConn, err := gorm.Open(configs.DB_DRIVER,
		fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			configs.DB_USER, configs.DB_PASSWORD, configs.DB_HOST, dbName))

	if err != nil {
		fmt.Println("Opt! Db connection error. Check configs/Config.go or copy Config.example to Config.go")
		os.Exit(1)
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
