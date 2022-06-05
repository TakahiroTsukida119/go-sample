package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

type SQLHandler struct {
	DB  *gorm.DB
	Err error
}

var dbConn *SQLHandler

func DBOpen() *SQLHandler {
	dbConn = NewSQLHandler()
	return dbConn
}

func DBClose() {
	sqlDB, _ := dbConn.DB.DB()
	if err := sqlDB.Close(); err != nil {
		return
	}
}

func NewSQLHandler() *SQLHandler {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	fmt.Println("==========================================================")
	fmt.Println("DB", user, password, host, port, dbName)
	fmt.Println("==========================================================")

	var db *gorm.DB
	var err error
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true&loc=Asia%2FTokyo"
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	//コネクションプールの最大接続数を設定。
	sqlDB.SetMaxIdleConns(100)
	//接続の最大数を設定。 nに0以下の値を設定で、接続数は無制限。
	sqlDB.SetMaxOpenConns(100)
	//接続の再利用が可能な時間を設定。dに0以下の値を設定で、ずっと再利用可能。
	sqlDB.SetConnMaxLifetime(100 * time.Second)

	sqlHandler := new(SQLHandler)
	db.Logger.LogMode(4)
	sqlHandler.DB = db

	return sqlHandler
}
