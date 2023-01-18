package auth

import (
	"FindMyDosen/database"
	"FindMyDosen/model/dto"
	"FindMyDosen/model/entity"
	"FindMyDosen/repository/auth_repo"
	"FindMyDosen/repository/redis_repo"
	"math/rand"
	"time"
)

func performRefreshToken(uid uint, refreshToken string) (dto.AuthDTO, error) {
	db := database.GetDB()
	var refreshRef entity.RefreshToken
	//if err := db.Preload("User").First(&refreshRef, "user_id = ?", uid).Error; err != nil {
	if err := db.Joins("lEFT JOIN users on refresh_tokens.user_id = users.id").First(&refreshRef, "user_id = ?", uid).Error; err != nil {
		return dto.AuthDTO{}, err
	}
	// check token
	err := checkPassword(refreshToken, refreshRef.RefreshKey)
	if err != nil {
		return dto.AuthDTO{}, err
	}

	token, err := generateToken(uid, true) //refreshRef.User.IsVerified)
	refresh, stored, err := generateRefreshKey()
	if err != nil {
		return dto.AuthDTO{}, err
	}
	err = db.Model(entity.RefreshToken{}).Where("user_id = ?", uid).Updates(
		entity.RefreshToken{
			RefreshKey: stored,
		}).Error
	if err != nil {
		return dto.AuthDTO{}, err
	}
	return dto.AuthDTO{
		Token:      token,
		RefreshKey: refresh,
	}, err
}

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
	token, err := generateToken(user.ID, user.IsVerified)
	refresh, stored, err := generateRefreshKey()
	if err != nil {
		return dto.AuthDTO{}, err
	}
	err = db.Model(entity.RefreshToken{}).Where("user_id = ?", user.ID).Updates(
		entity.RefreshToken{
			RefreshKey: stored,
		}).Error
	if err != nil {
		return dto.AuthDTO{}, err
	}
	return dto.AuthDTO{
		Token:      token,
		RefreshKey: refresh,
	}, err
}

func performUserRegistration(user *dto.NewUserDTO) (error, dto.AuthDTO) {
	db := database.GetDB()
	hashed, err := auth_repo.HashPassword(user.Password)
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

	t, err := generateToken(newUser.ID, newUser.IsVerified)
	if err != nil {
		return err, dto.AuthDTO{}
	}
	refresh, err := redis_repo.NewRefreshToken(newUser.ID)

	//refresh, stored, err := generateRefreshKey()
	//storedRefresh := entity.RefreshToken{
	//	UserID:     newUser.ID,
	//	RefreshKey: stored,
	//}
	//if err = db.Create(&storedRefresh).Error; err != nil {
	//	return err, dto.AuthDTO{}
	//}
	if err != nil {
		return err, dto.AuthDTO{}
	}
	return nil, dto.AuthDTO{
		Token:      t,
		RefreshKey: refresh,
	}
}

func generateRefreshKey() (string, string, error) {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	length := rand.Intn(15-8) + 8 // pick a random length between 8 and 15
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	randomKey := string(b)
	hashed, err := auth_repo.HashPassword(randomKey)
	return randomKey, hashed, err
}
