package repositories

import (
	"Demo/models"
	"database/sql"
	"log"
)

type IUserRepository interface {
	CreateUser(userAuth models.UserAuth) error
	GetUserById(id int, userStored *models.UserStored) error
	GetUserByEmail(email string, userStored *models.UserStored) error
}

type userRepository struct {
	db *sql.DB
}

func InitUserRepository(db *sql.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(userAuth models.UserAuth) error {
	if _, e := ur.db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", userAuth.Email, userAuth.Password); e != nil {
		log.Println("[error] creating user by email in repository: " + e.Error())
		return e
	}
	return nil
}

func (ur *userRepository) GetUserById(id int, userStored *models.UserStored) error {
	r := ur.db.QueryRow("SELECT id, email, password FROM public.users WHERE id = $1", id)
	if e := r.Scan(&userStored.Id, &userStored.Email, &userStored.Password /*, &userStored.CreatedAt, &userStored.UpdatedAt*/); e != nil {
		log.Println("[error] getting user by id in repository: " + e.Error())
		return e
	}
	return nil
}

func (ur *userRepository) GetUserByEmail(email string, userStored *models.UserStored) error {
	r := ur.db.QueryRow("SELECT id, email, password FROM public.users WHERE email = $1", email)
	if e := r.Scan(&userStored.Id, &userStored.Email, &userStored.Password /*, &userStored.CreatedAt, &userStored.UpdatedAt*/); e != nil {
		log.Println("[error] getting user in repository: " + e.Error())
		return e
	}
	return nil
}
