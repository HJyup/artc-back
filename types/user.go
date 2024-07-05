package types

import "time"

type Speciality string

const (
	Musician     Speciality = "musician"
	Actor        Speciality = "actor"
	VisualArtist Speciality = "visual_artist"
	Writer       Speciality = "writer"
	Designer     Speciality = "designer"
	Dancer       Speciality = "dancer"
	Photographer Speciality = "photographer"
	Filmmaker    Speciality = "filmmaker"
)

type User struct {
	ID         int        `json:"id"`
	Avatar     string     `json:"avatar"`
	FirstName  string     `json:"first_name"`
	SecondName string     `json:"second_name"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Speciality Speciality `json:"speciality"`
	Location   string     `json:"location"`
	IsAccepted bool       `json:"is_accepted"`
	IsReviewer bool       `json:"is_reviewer"`
	CreatedAt  time.Time  `json:"created_at"`
}

type RegisterUserPayload struct {
	FirstName  string `json:"first_name" validate:"required"`
	SecondName string `json:"second_name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6,max=32"`
	Speciality string `json:"speciality" validate:"required"`
	Location   string `json:"location" validate:"required"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
