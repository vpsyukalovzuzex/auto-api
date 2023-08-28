package usecases

import (
	"Demo/models"
	"Demo/repositories"
	"Demo/utils"
	"Demo/validators"

	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(userAuth models.UserAuth) error
	SignIn(userAuth models.UserAuth, userStored *models.UserStored) error
	GetUserById(id int, userStored *models.UserStored) error
}

type userUsecase struct {
	ur repositories.IUserRepository
	uv validators.IUserValidator
}

func InitUserUsecase(ur repositories.IUserRepository, uv validators.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(userAuth models.UserAuth) error {
	if e := uu.uv.Validate(userAuth); e != nil {
		return &utils.CustomError{Code: 400, Description: "Invalid parameters"}
	}
	userStored := models.UserStored{}
	uu.ur.GetUserByEmail(userAuth.Email, &userStored)
	if userAuth.Email == userStored.Email {
		return &utils.CustomError{Code: 400, Description: "User already exists"}
	}
	hash, e := bcrypt.GenerateFromPassword([]byte(userAuth.Password), bcrypt.DefaultCost)
	if e != nil {
		return &utils.CustomError{Code: 500, Description: e.Error()}
	}
	userAuth.Password = string(hash)
	return uu.ur.CreateUser(userAuth)
}

func (uu *userUsecase) SignIn(userAuth models.UserAuth, userStored *models.UserStored) error {
	if e := uu.uv.Validate(userAuth); e != nil {
		return &utils.CustomError{Code: 400, Description: "Invalid parameters"}
	}
	if e := uu.ur.GetUserByEmail(userAuth.Email, userStored); e != nil {
		return &utils.CustomError{Code: 400, Description: "User doesn't exists"}
	}
	if e := bcrypt.CompareHashAndPassword([]byte(userStored.Password), []byte(userAuth.Password)); e != nil {
		return &utils.CustomError{Code: 400, Description: "Wrong password"}
	}
	return nil
}

func (uu *userUsecase) GetUserById(id int, userStored *models.UserStored) error {
	return uu.ur.GetUserById(id, userStored)
}
