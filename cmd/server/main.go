package main

import (
	"c0fee-api/controller"
	"c0fee-api/infrastructure/db"
	"c0fee-api/infrastructure/s3"
	"c0fee-api/repository"
	"c0fee-api/router"
	"c0fee-api/usecase"
	"log"

	"github.com/go-playground/validator"
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
	beanRatingRepository := repository.NewBeanRatingRepository(db)
	countryRepository := repository.NewCountryRepository(db)
	farmRepository := repository.NewFarmRepository(db)
	processMethodRepository := repository.NewProcessMethodRepository(db)
	roasterRepository := repository.NewRoasterRepository(db)
	userRepository := repository.NewUserRepository(db)
	varietyRepository := repository.NewVarietyRepository(db)

	// validator
	validator := validator.New()

	// usecases
	areaUsecase := usecase.NewAreaUsecase(areaRepository)
	beanUseCase := usecase.NewBeanUsecase(userRepository, beanRepository, beanRatingRepository, s3Service, validator)
	countryUsecase := usecase.NewCountryUsecase(countryRepository)
	farmUsecase := usecase.NewFarmUsecase(farmRepository)
	processMethodUsecase := usecase.NewProcessMethodUsecase(processMethodRepository)
	roasterUsecase := usecase.NewRoasterUsecase(roasterRepository)
	roastLevelUsecase := usecase.NewRoastLevelUsecase()
	userUseCase := usecase.NewUserUsecase(userRepository, beanRepository, s3Service)
	varietyUsecase := usecase.NewVarietyUsecase(varietyRepository)

	// controllers
	areaController := controller.NewAreaController(areaUsecase)
	beanController := controller.NewBeanController(beanUseCase)
	countryController := controller.NewCountryController(countryUsecase)
	farmController := controller.NewFarmController(farmUsecase)
	processMethodController := controller.NewProcessMethodController(processMethodUsecase)
	roasterController := controller.NewRoasterController(roasterUsecase)
	roastLevelController := controller.NewRoastLevelController(roastLevelUsecase)
	userController := controller.NewUserController(userUseCase)
	varietyController := controller.NewVarietyController(varietyUsecase)

	e := router.NewRouter(userController, beanController, countryController, roasterController, areaController, farmController, varietyController, processMethodController, roastLevelController)
	e.Logger.Fatal(e.Start(":8080"))
}
