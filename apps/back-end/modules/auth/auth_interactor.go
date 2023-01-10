package auth

import (
	"FindMyDosen/config"
	"FindMyDosen/database"
	"FindMyDosen/model/dto"
	"FindMyDosen/model/entity"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
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
	hashed, err := hashPassword(user.Password)
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
		BirthDate:    user.BirthDate,
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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPassword(password string, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err
}

func generateToken(uid uint) (string, error) {
	claims := entity.JwtClaims{
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetJWTSecret()))
}
