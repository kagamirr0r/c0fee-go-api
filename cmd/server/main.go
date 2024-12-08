package main

import (
	"c0fee-api/controller"
	"c0fee-api/db"
	"c0fee-api/repository"
	"c0fee-api/router"
	"c0fee-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUseCase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
