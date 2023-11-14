package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dsn := "host=127.0.0.1 user=" + dbUser + " password=" + dbPassword + " dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.Migrator().DropTable(&User{})

	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}

	fmt.Println("*******DB READY*******")

	DB = db
}
