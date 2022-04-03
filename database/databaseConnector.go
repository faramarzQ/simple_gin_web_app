package database

import (
	"gin/config"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func init() {
	var err error

	dsn := makeDataSourceName()

	Connection, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: config.DB("DRIVER_NAME"),
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

// data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
func makeDataSourceName() string {
	dsn := config.DB("USERNAME")

	if config.DB("PASSWORD") != "" {
		dsn += config.DB("PASSWORD")
	}

	dsn += "@tcp(" + config.DB("HOST") + ":" + config.DB("PORT") + ")"
	dsn += "/" + config.DB("NAME")

	dsn += "?charset=utf8&parseTime=True&loc=Local"

	return dsn
}

func Con() *gorm.DB {
	return Connection
}
