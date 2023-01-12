package auth

import (
	"FindMyDosen/database"
	"FindMyDosen/model/dto"
	"FindMyDosen/model/entity"
)

func performUserLogin(userData *dto.LoginUserDTO) (dto.AuthDTO, error) {
	db := database.GetDB()
	var user entity.User
	err := db.Where("email = ?", userData.Email).First(&user).Error
	if err != nil {
		return dto.AuthDTO{}, err
	}
	err = checkPassword(userData.Password, user.Password)
	if err != nil {
		return dto.AuthDTO{}, err
	}
	token, err := generateToken(user.ID)
	return dto.AuthDTO{
		Token:      token,
		RefreshKey: "",
	}, err
}

func performUserRegistration(user *dto.NewUserDTO) (error, dto.AuthDTO) {
	db := database.GetDB()
	hashed, err := HashPassword(user.Password)
	if err != nil {
		return err, dto.AuthDTO{}
	}
	newUser := entity.User{
		Email:        user.Email,
		Username:     user.Username,
		FullName:     user.FullName,
		Password:     hashed,
		Avatar:       user.Avatar,
		UniversityID: user.UniversityID,
		Npm:          user.Npm,
		BirthDate:    user.BirthDate.Time,
	}
	if err = db.Create(&newUser).Error; err != nil {
		return err, dto.AuthDTO{}
	}

	t, err := generateToken(newUser.ID)
	if err != nil {
		return err, dto.AuthDTO{}
	}
	return nil, dto.AuthDTO{
		Token:      t,
		RefreshKey: "Ini refresh",
	}
}
