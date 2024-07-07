package types

import (
	"database/sql"
	"time"
)

type Speciality string

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id string) (*User, error)
	CreateUser(user User) error
}

type User struct {
	ID           string         `json:"id"`
	AvatarURL    sql.NullString `json:"avatar"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
	SpecialityID int            `json:"speciality_id"`
	Location     string         `json:"location"`
	IsAccepted   bool           `json:"is_accepted"`
	IsReviewer   bool           `json:"is_reviewer"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type RegisterUserPayload struct {
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=6,max=32"`
	SpecialityID int    `json:"speciality_id"`
	Location     string `json:"location" validate:"required"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
