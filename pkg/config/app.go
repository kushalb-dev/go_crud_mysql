package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	godotenv.Load()
	uname := os.Getenv("UNAME")
	password := os.Getenv("PASSWORD")
	address := os.Getenv("ADDRESS")
	dbname := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", uname, password, address, dbname)
	fmt.Println(dsn)
	d, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
