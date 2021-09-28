package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_const "handlegeo/const"
	"log"
	"os"
	"time"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func getDBLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   	// Slow SQL threshold
			LogLevel:                   logger.Info, 	// Log level
			IgnoreRecordNotFoundError: 	true,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  	true,          	// Disable color
		},
	)
	return newLogger
}

func Conn() (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dbLogger := getDBLogger()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		_const.MYSQLUSER, _const.MYSQLPASS, _const.MYSQLHOST,_const.MYSQLDB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dbLogger})
	if err != nil {
		return nil, err
	}
	//migrateDB(db)
	DB = db
	return DB, nil
}
