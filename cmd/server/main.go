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

	// users
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUseCase)

	// beans
	beanRepository := repository.NewBeanRepository(db)
	beanUseCase := usecase.NewBeanUsecase(userRepository, beanRepository)
	beanController := controller.NewBeanController(beanUseCase)

	e := router.NewRouter(userController, beanController)
	e.Logger.Fatal(e.Start(":8080"))
}
