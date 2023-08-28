package models

type UserCar struct {
	Id     int    `json:"id" db:"user_car_id"`
	Number string `json:"number" db:"user_car_number"`
	Car    Car    `json:"car"`
	// CreatedAt time.Time `json:"created_at" db:"user_car_created_at"`
	// UpdatedAt time.Time `json:"updated_at" db:"user_car_updated_at"`
}
