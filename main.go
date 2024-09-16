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

	pr := repository.NewProjectRepository(db)
	ps := service.NewProjectService(pr)
	pc := controller.NewProjectController(ps)
	e := router.NewRouter(uc, pc)

	err := e.Start(":8080")

	if err != nil {
		e.Logger.Fatal(err)
	}

}
