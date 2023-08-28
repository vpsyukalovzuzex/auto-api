package models

type Car struct {
	Id     int    `json:"id" db:"car_id"`
	Name   string `json:"name" db:"car_name"`
	Engine Engine `json:"engine"`
	// CreatedAt time.Time `json:"created_at" db:"car_created_at"`
	// UpdatedAt time.Time `json:"updated_at" db:"car_updated_at"`
}
