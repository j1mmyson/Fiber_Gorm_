package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/j1mmyson/hackathon_backend/config"
	"github.com/j1mmyson/hackathon_backend/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Wrong")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect DB.")
	}

	fmt.Println("Success to connect DB.")

	DB.AutoMigrate(&model.User{}, &model.Article{}, &model.UserInfo{}, &model.Tag{}, &model.UserTag{}, &model.ArticleTag{})
	fmt.Println("DB Migrated.")
}
