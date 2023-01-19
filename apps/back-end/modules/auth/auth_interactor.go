package auth

import (
	"FindMyDosen/database"
	"FindMyDosen/model/dto"
	"FindMyDosen/model/entity"
	"FindMyDosen/repository/auth_repo"
	"FindMyDosen/repository/redis_repo"
)

type RefreshError struct{}

func (m *RefreshError) Error() string {
	return "Error Validating Refresh Key"
}

func performRefreshToken(uid uint, refreshToken string) (dto.AuthDTO, error) {
	key, err := redis_repo.GetRefreshToken(uid)
	if err != nil {
		return dto.AuthDTO{}, err
	}
	if key != refreshToken {
		return dto.AuthDTO{}, &RefreshError{}
	}

	token, err := generateToken(uid, true) //refreshRef.User.IsVerified)
	refresh, err := redis_repo.NewRefreshToken(uid)
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
		println("USer ", user.Username)
		//return dto.AuthDTO{}, err
	}
	err = checkPassword(userData.Password, user.Password)
	if err != nil {
		return dto.AuthDTO{}, err
	}
	token, err := generateToken(user.ID, user.IsVerified)
	refresh, err := redis_repo.GetRefreshToken(user.ID)
	if err != nil {
		refresh, err = redis_repo.NewRefreshToken(user.ID)
	}
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
	if err != nil {
		return err, dto.AuthDTO{}
	}
	return nil, dto.AuthDTO{
		Token:      t,
		RefreshKey: refresh,
	}
}
