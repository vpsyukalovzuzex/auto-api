package usecases

import (
	"Demo/models"
	"Demo/repositories"
)

type ICarUsecase interface {
	GetAllUserCarsByUserId(userId int, cars *[]models.UserCar) error
	GetAllUserEnginesByUserId(userId int, engines *[]models.Engine) error
	GetCarEngine(carId int, engine *models.Engine) error
}
type carUsecase struct {
	cr repositories.ICarRepository
}

func InitCarUsecase(cr repositories.ICarRepository) ICarUsecase {
	return &carUsecase{cr}
}

func (cu *carUsecase) GetAllUserCarsByUserId(userId int, cars *[]models.UserCar) error {
	return cu.cr.GetAllUserCarsByUserId(userId, cars)
}

func (cu *carUsecase) GetAllUserEnginesByUserId(userId int, engines *[]models.Engine) error {
	return cu.cr.GetAllUserEnginesByUserId(userId, engines)
}

func (cu *carUsecase) GetCarEngine(carId int, engine *models.Engine) error {
	return cu.cr.GetCarEngine(carId, engine)
}
