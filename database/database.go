package database

import (
	"fmt"
	"uts/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn:= "root@tcp(localhost)/awikwok?charset=utf8mb4&parseTime=True&loc=Local"
	
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal Bang :(")
	}
	fmt.Println("Sukes :)")
	
	DB = conn
	DB.AutoMigrate(models.User{} )
}
