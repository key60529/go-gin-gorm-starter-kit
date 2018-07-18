package database

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// Register mysql driver for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

const (
	//DbHost ...
	DbHost = ""
	//DbUser ...
	DbUser = ""
	//DbPassword ...
	DbPassword = ""
	//DbName ...
	DbName = ""
)

func init() {
	dbinfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DbUser, DbPassword, DbHost, DbName)

	db, err = ConnectDB(dbinfo)
	if err != nil {
		fmt.Println(err)
	}
}

// ConnectDB open a connection
func ConnectDB(dataSourceName string) (*gorm.DB, error) {
	db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("failed to connect database")
	}
	return db, nil
}

//GetDB return database
func GetDB() *gorm.DB {
	return db
}