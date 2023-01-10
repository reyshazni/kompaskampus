package dto

type NewUserDTO struct {
	Email      string `json:"email" validate:"required,email"`
	Username   string `json:"username" validate:"required"`
	FullName   string `json:"fullName" validate:"required"`
	Password   string `json:"password" validate:"required"`
	RePassword string `json:"re_password" validate:"required"`
	Avatar     string `json:"avatar" validate:"omitempty,required"`
	// BirthDate time.Time `json:"birthDate"`
	University string `json:"university" validate:"required"`
	Npm        string `json:"npm" validate:"required"`
}
