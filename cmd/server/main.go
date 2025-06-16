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
	areaRepository := repository.NewAreaRepository(db)
	beanRepository := repository.NewBeanRepository(db)
	countryRepository := repository.NewCountryRepository(db)
	farmRepository := repository.NewFarmRepository(db)
	roasterRepository := repository.NewRoasterRepository(db)
	userRepository := repository.NewUserRepository(db)

	// usecases
	areaUsecase := usecase.NewAreaUsecase(areaRepository)
	beanUseCase := usecase.NewBeanUsecase(userRepository, beanRepository, s3Service)
	countryUsecase := usecase.NewCountryUsecase(countryRepository)
	farmUsecase := usecase.NewFarmUsecase(farmRepository)
	roasterUsecase := usecase.NewRoasterUsecase(roasterRepository)
	userUseCase := usecase.NewUserUsecase(userRepository, beanRepository, s3Service)

	// controllers
	areaController := controller.NewAreaController(areaUsecase)
	beanController := controller.NewBeanController(beanUseCase)
	countryController := controller.NewCountryController(countryUsecase)
	farmController := controller.NewFarmController(farmUsecase)
	roasterController := controller.NewRoasterController(roasterUsecase)
	userController := controller.NewUserController(userUseCase)

	e := router.NewRouter(userController, beanController, countryController, roasterController, areaController, farmController)
	e.Logger.Fatal(e.Start(":8080"))
}
