package models

type UserAuth struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6"`
}

type UserStored struct {
	Id       int
	Email    string
	Password string
	// CreatedAt string
	// UpdatedAt string
}

type User struct {
	Id    int    `json:"id" db:"user_id"`
	Email string `json:"email" db:"user_email"`
	// CreatedAt string `json:"created_at" db:"user_created_at"`
	// UpdatedAt string `json:"updated_at" db:"user_updated_at"`
}
