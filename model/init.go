package model

import (
	"blog/model/blogs"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("postgre.host"), viper.GetInt("postgre.port"), viper.GetString("postgre.username"),
		viper.GetString("postgre.password"), viper.GetString("postgre.dbname"))
	fmt.Println(psqlInfo)
	db, err := gorm.Open("postgres", psqlInfo)
	defer db.Close()
	if err != nil {
		//fmt.Println("err when open conn",err)
		panic(err)
	}
	_ = createTable(db)
	fmt.Println("connect successfully!")
}

func createTable(db *gorm.DB) error {

	if err := db.AutoMigrate(&blogs.User{}).Error; err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&blogs.Contact{}).Error; err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&blogs.Blog{}).Error; err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&blogs.Message{}).Error; err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&blogs.Comment{}).Error; err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&blogs.Subscribe{}).Error; err != nil {
		panic(err)
	}
	return nil
}

func GetDB() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("postgre.host"), viper.GetInt("postgre.port"), viper.GetString("postgre.username"),
		viper.GetString("postgre.password"), viper.GetString("postgre.dbname"))
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}

	return db
}
