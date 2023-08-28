package models

type Engine struct {
	Id    int    `json:"id" db:"engine_id"`
	Name  string `json:"name" db:"engine_name"`
	Power int    `json:"power" db:"engine_power"`
	// CreatedAt time.Time `json:"created_at" db:"engine_created_at"`
	// UpdatedAt time.Time `json:"updated_at" db:"engine_updated_at"`
}
