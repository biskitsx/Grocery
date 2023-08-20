package database

import (
	"errors"

	"github.com/biskitsx/Grocery/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func StartMySQL() error {
	var err error
	dsn := "root:root@tcp(localhost:8889)/Grocery?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("can't connect to database")
	}

	Db.AutoMigrate(&model.Category{}, &model.Product{}, &model.User{})
	return nil
}
