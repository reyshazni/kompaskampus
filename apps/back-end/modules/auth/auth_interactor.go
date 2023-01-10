package auth

import (
	"FindMyDosen/database"
	"FindMyDosen/model/dto"
	"FindMyDosen/model/entity"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func performUserRegistration(user *dto.NewUserDTO) (error, dto.AuthDTO) {
	db := database.GetDB()
	hashed, err := hashPassword(user.Password)
	if err != nil {
		return err, dto.AuthDTO{}
	}
	newUser := entity.User{
		Email:      user.Email,
		Username:   user.Username,
		FullName:   user.FullName,
		Password:   hashed,
		Avatar:     user.Avatar,
		University: user.University,
		Npm:        user.Npm,
	}
	if err = db.Create(&newUser).Error; err != nil {
		return err, dto.AuthDTO{}
	}

	claims := entity.JwtClaims{
		newUser.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
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

func checkPassword(password string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
