package utils

import (
	"fmt"

	"github.com/awesomexu/go-realworld/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(){
	db, err := gorm.Open(mysql.Open(viper.GetString("mysqlDNS")), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
	fmt.Println("数据库连接成功")
	DB = db

	db.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Comment{},
		&models.Tag{},
	)
}