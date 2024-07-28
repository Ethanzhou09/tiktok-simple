package db

import (
	"fmt"
	"log"
	"tiktok-simple/pkg/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(){
	viper.Init()
	dbconfig := viper.GetdbConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconfig.Username, dbconfig.Password, dbconfig.Host, dbconfig.Port, dbconfig.Database)
	DB,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&User{},&Video{},&Comment{},&FavoriteCommentRelation{},&Message{},&FavoriteVideoRelation{},&FollowRelation{})
}

func GetDB() *gorm.DB {
	return DB
}