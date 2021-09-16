package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_const "handlegeo/const"
	"handlegeo/utils"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func DBConn() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		_const.MYSQLUSER, _const.MYSQLPASS, _const.MYSQLDB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	utils.CheckErr(err)
	DB = db
	return DB
}

func MigrateDB(db *gorm.DB) {
	// DB를
	err := db.AutoMigrate(&Cell{})
	utils.CheckErr(err)
}

func GetDB() *gorm.DB {
	return DB
}
