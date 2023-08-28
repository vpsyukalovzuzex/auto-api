package repositories

import (
	"Demo/models"
	"Demo/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/jackskj/carta"
)

type ICarRepository interface {
	GetAllUserCarsByUserId(userId int, cars *[]models.UserCar) error
	GetAllUserEnginesByUserId(userId int, cars *[]models.Engine) error
	GetCarEngine(carId int, engine *models.Engine) error
}

type carRepository struct {
	db *sql.DB
}

func InitCarRepository(db *sql.DB) ICarRepository {
	return &carRepository{db}
}

func (cr *carRepository) GetAllUserCarsByUserId(userId int, cars *[]models.UserCar) error {
	q := `SELECT
	cars.id as car_id,
	cars.name as car_name,
	cars.created_at as car_created_at,
	cars.updated_at as car_updated_at,
	users_cars.id as user_car_id,
	users_cars.number as user_car_number,
	users_cars.created_at as user_car_created_at,
	users_cars.updated_at as user_car_updated_at,
	engines.id as engine_id,
	engines.name as engine_name,
	engines.power as engine_power,
	engines.created_at as engine_created_at,
	engines.updated_at as engine_updated_at
	FROM users_cars
	INNER JOIN users ON users_cars.id_user = users.id
	INNER JOIN cars ON users_cars.id_car = cars.id
	INNER JOIN engines ON cars.id_engine = engines.id
	WHERE users_cars.id_user = $1`
	r, e := cr.db.Query(q, userId)
	if e != nil {
		return e
	}
	if e := carta.Map(r, cars); e != nil {
		return e
	}
	return nil
}

func (cr *carRepository) GetAllUserEnginesByUserId(userId int, cars *[]models.Engine) error {
	q := `SELECT
	DISTINCT engines.id AS engine_id,
	engines.name AS engine_name,
	engines.power AS engine_power
	FROM users_cars
	INNER JOIN cars ON cars.id = users_cars.id_car
	INNER JOIN engines ON engines.id = cars.id_engine
	WHERE id_user = $1`
	r, e := cr.db.Query(q, userId)
	if e != nil {
		return e
	}
	if e := carta.Map(r, cars); e != nil {
		return e
	}
	return nil
}

func (cr *carRepository) GetCarEngine(carId int, engine *models.Engine) error {
	q := `select
	distinct engines.id as engine_id,
	engines.name as engine_name,
	engines.power as engine_power
	from cars
	inner join engines on cars.id_engine = engines.id
	where cars.id = $1`
	r := cr.db.QueryRow(q, carId)
	if e := r.Err(); e != nil {
		return e
	}
	if e := r.Scan(&engine.Id, &engine.Name, &engine.Power); e != nil {
		if e == sql.ErrNoRows {
			return &utils.CustomError{
				Code:        http.StatusNotFound,
				Description: fmt.Sprintf("Can't find engine for car id = %d", carId),
			}
		}
		return e
	}
	return nil
}
