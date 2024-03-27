package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("CURRENT_DB_USER"),
		os.Getenv("CURRENT_DB_PW"),
		os.Getenv("CURRENT_DB_HOST"),
		os.Getenv("CURRENT_DB_PORT"),
		os.Getenv("CURRENT_DB_NAME"),
	)

	fmt.Println(url)

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
