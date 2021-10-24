package database

import (
	"fmt"

	"github.com/abhishekratnam/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	user := "user1"
	password := ""
	host := "localhost"
	port := "3306"
	dbname := "sys"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("v:Could not connect to database")
	}
	DB = conn
	//Migrate
	conn.AutoMigrate(&models.User{})
}
