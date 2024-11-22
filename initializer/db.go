package initializer

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dburl := os.Getenv("DATABASE_URL")
	var Error error
	DB, Error = gorm.Open(mysql.Open(dburl), &gorm.Config{})
	if Error != nil {
		log.Fatal("Error in connecting database", Error.Error())
	}
	fmt.Println("🚀 Database connected successfully !!!")
}
