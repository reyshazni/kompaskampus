package dto

import (
	"time"
)

type BirthDate struct {
	Time time.Time
}

type NewUserDTO struct {
	Email        string    `json:"email" validate:"required,email"`
	Username     string    `json:"username" validate:"required"`
	FullName     string    `json:"fullName" validate:"required"`
	Password     string    `json:"password" validate:"required"`
	RePassword   string    `json:"re_password" validate:"required"`
	Avatar       string    `json:"avatar" validate:"omitempty,required"`
	BirthDate    BirthDate `json:"birth_date" validate:"required"`
	UniversityID int       `json:"university_id" validate:"required"`
	Npm          string    `json:"npm" validate:"required"`
}

func (mt *BirthDate) UnmarshalJSON(b []byte) error {
	// Define a custom format to parse
	const format = "2006-01-02"

	// Extract the string from b
	s := string(b)
	// Remove the surrounding quotes
	s = s[1 : len(s)-1]
	// Parse the string using the custom format
	newTime, err := time.Parse(format, s)
	if err != nil {
		return err
	}
	mt.Time = newTime
	return nil
}

type LoginUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
