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
	s3Service, err := s3.NewS3Service()
	if err != nil {
		log.Fatal(err)
	}

	// users
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository, s3Service)
	userController := controller.NewUserController(userUseCase)

	// beans
	beanRepository := repository.NewBeanRepository(db)
	beanUseCase := usecase.NewBeanUsecase(userRepository, beanRepository, s3Service)
	beanController := controller.NewBeanController(beanUseCase)

	// countries
	countryRepository := repository.NewCountryRepository(db)
	countryUsecase := usecase.NewCountryUsecase(countryRepository)
	countryController := controller.NewCountryController(countryUsecase)

	// roasters
	roasterRepository := repository.NewRoasterRepository(db)
	roasterUsecase := usecase.NewRoasterUsecase(roasterRepository)
	roasterController := controller.NewRoasterController(roasterUsecase)

	e := router.NewRouter(userController, beanController, countryController, roasterController)
	e.Logger.Fatal(e.Start(":8080"))
}
