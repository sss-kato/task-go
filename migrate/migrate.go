package main

import (
	"fmt"
	"task-go/db"
	"task-go/model"
)

func main() {
	con := db.NewDB()
	defer fmt.Println("Migration Success")
	defer db.CloseDB(con)
	con.AutoMigrate(&model.User{})
}
