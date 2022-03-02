package database

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func init() {
	var err error
	Connection, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        "root@tcp(127.0.0.1:3306)/go_db?charset=utf8&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

func Con() *gorm.DB {
	return Connection
}
