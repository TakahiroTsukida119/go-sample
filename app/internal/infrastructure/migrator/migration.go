package migrator

import (
	"github.com/TakahiroTsukida119/go-sample.git/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Run() {
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_DATABASE")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to migration")
	}
}
