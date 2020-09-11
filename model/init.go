package model
import (
	"blog/model/blogs"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Init(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("postgre.host"),viper.GetInt("postgre.port"),viper.GetString("postgre.username"),
		viper.GetString("postgre.password"),viper.GetString("postgre.dbname"))
	fmt.Println(psqlInfo)
	db,err:=gorm.Open("postgres",psqlInfo)
	defer db.Close()
	if err != nil {
		//fmt.Println("err when open conn",err)
		panic(err)
	}
	_=createTable(db)
	fmt.Println("connect successfully!")
}

func createTable(db *gorm.DB) error {
	if !db.HasTable(&blogs.Account{}){
		if err:=db.CreateTable(&blogs.Account{}).Error;err!=nil{
			panic(err)
		}
	}
	if !db.HasTable(&blogs.Contact{}){
		if err:=db.CreateTable(&blogs.Contact{}).Error;err!=nil{
			panic(err)
		}
	}
	if !db.HasTable(&blogs.Author{}){
		if err:=db.CreateTable(&blogs.Author{}).Error;err!=nil{
			panic(err)
		}
	}
	if !db.HasTable(&blogs.Blog{}){
		if err:=db.CreateTable(&blogs.Blog{}).Error;err!=nil{
			panic(err)
		}
	}
	if !db.HasTable(&blogs.Message{}){
		if err:=db.CreateTable(&blogs.Message{}).Error;err!=nil{
			panic(err)
		}
	}
	if !db.HasTable(&blogs.Comment{}){
		if err:=db.CreateTable(&blogs.Comment{}).Error;err!=nil{
			panic(err)
		}
	}
	return nil
}

func GetDB() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("postgre.host"),viper.GetInt("postgre.port"),viper.GetString("postgre.username"),
		viper.GetString("postgre.password"),viper.GetString("postgre.dbname"))
	db,err:=gorm.Open("postgres",psqlInfo)
	if err != nil {
		fmt.Println(err)
	}

	return db
}