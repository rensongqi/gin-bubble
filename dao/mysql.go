package dao

import (
	"gin-bubble/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:123456@(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err!= nil{
		return
	}
	return DB.DB().Ping()
}

func InitModel()  {
	DB.AutoMigrate(&models.Todo{})
}

func Close() {
	DB.Close()
}