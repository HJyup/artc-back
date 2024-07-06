package types

import (
	"github.com/google/uuid"
	"time"
)

type Speciality string

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id string) (*User, error)
	CreateUser(user User) error
}

type User struct {
	PK           int       `json:"pk"`
	ID           uuid.UUID `json:"id"`
	Avatar       string    `json:"avatar"`
	FirstName    string    `json:"first_name"`
	SecondName   string    `json:"second_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	SpecialityID int       `json:"speciality_id"`
	Location     string    `json:"location"`
	IsAccepted   bool      `json:"is_accepted"`
	IsReviewer   bool      `json:"is_reviewer"`
	CreatedAt    time.Time `json:"created_at"`
}

type RegisterUserPayload struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name" validate:"required"`
	SecondName string    `json:"second_name" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required,min=6,max=32"`
	Speciality string    `json:"speciality" validate:"required"`
	Location   string    `json:"location" validate:"required"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
