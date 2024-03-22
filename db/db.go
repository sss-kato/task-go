package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {

	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("CURRENT_DB_USER"),
		os.Getenv("CURRENT_DB_PW"),
		os.Getenv("CURRENT_DB_HOST"),
		os.Getenv("CURRENT_DB_PORT"),
		os.Getenv("CURRENT_DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Success")
	return db

}

func CloseDB(db *gorm.DB) {
	con, _ := db.DB()
	if err := con.Close(); err != nil {
		log.Fatal(err)
	}
}
