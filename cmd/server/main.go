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
	// initialize infrastructure instances
	db := db.NewDB()
	s3Service, err := s3.NewS3Service()
	if err != nil {
		log.Fatal(err)
	}

	//repositories
	beanRepository := repository.NewBeanRepository(db)
	countryRepository := repository.NewCountryRepository(db)
	roasterRepository := repository.NewRoasterRepository(db)
	userRepository := repository.NewUserRepository(db)

	// usecases
	beanUseCase := usecase.NewBeanUsecase(userRepository, beanRepository, s3Service)
	countryUsecase := usecase.NewCountryUsecase(countryRepository)
	roasterUsecase := usecase.NewRoasterUsecase(roasterRepository)
	userUseCase := usecase.NewUserUsecase(userRepository, beanRepository, s3Service)

	// controllers
	beanController := controller.NewBeanController(beanUseCase)
	countryController := controller.NewCountryController(countryUsecase)
	roasterController := controller.NewRoasterController(roasterUsecase)
	userController := controller.NewUserController(userUseCase)

	e := router.NewRouter(userController, beanController, countryController, roasterController)
	e.Logger.Fatal(e.Start(":8080"))
}
