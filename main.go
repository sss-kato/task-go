package main

import (
	"task-go/controller"
	"task-go/db"
	"task-go/repository"
	"task-go/router"
	"task-go/service"
)

func main() {
	db := db.NewDB()
	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur)
	uc := controller.NewUserController(us)
	e := router.NewRouter(uc)

	err := e.Start(":8080")

	if err != nil {
		e.Logger.Fatal(err)
	}

	slice4 := make([]string, 3)
}
