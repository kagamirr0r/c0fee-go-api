package main

import (
	"c0fee-api/controller"
	"c0fee-api/infrastructure/db"
	"c0fee-api/infrastructure/s3"
	"c0fee-api/repository"
	"c0fee-api/router"
	"c0fee-api/usecase"
	"log"
)

func main() {
	db := db.NewDB()
	s3Client, err := s3.NewS3Client()
	if err != nil {
		log.Fatal(err)
	}

	// users
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository, s3Client)
	userController := controller.NewUserController(userUseCase)

	// beans
	beanRepository := repository.NewBeanRepository(db)
	beanUseCase := usecase.NewBeanUsecase(userRepository, beanRepository)
	beanController := controller.NewBeanController(beanUseCase)

	e := router.NewRouter(userController, beanController)
	e.Logger.Fatal(e.Start(":8080"))
}
