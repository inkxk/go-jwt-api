package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func ConnectDb() {
	// connect database mysql
	dsn := "root:mysql123@tcp(127.0.0.1:3306)/TEST_DATABASE?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// auto migrate database
	Db.AutoMigrate(&User{})
}
