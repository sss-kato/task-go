package db

import (
	"fmt"
	"log"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PW"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
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

func CreateDBMock() (sqlmock.Sqlmock, *gorm.DB) {
	dbm, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// defer dbm.Close()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 dbm,
		PreferSimpleProtocol: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("an error '%s' was not expected when initializing a gorm db connection", err)
	}

	return mock, gormDB

}
