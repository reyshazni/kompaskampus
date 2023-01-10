package dto

import "time"

type NewUserDTO struct {
	Email        string    `json:"email" validate:"required,email"`
	Username     string    `json:"username" validate:"required"`
	FullName     string    `json:"fullName" validate:"required"`
	Password     string    `json:"password" validate:"required"`
	RePassword   string    `json:"re_password" validate:"required"`
	Avatar       string    `json:"avatar" validate:"omitempty,required"`
	BirthDate    time.Time `json:"birthDate" validate:"required"`
	UniversityID int       `json:"university_id" validate:"required"`
	Npm          string    `json:"npm" validate:"required"`
}

type LoginUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
