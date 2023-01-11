package user

import (
	"FindMyDosen/database"
	"FindMyDosen/model/entity"
	"FindMyDosen/modules/auth"
	"time"
)

type UserInfoDTO struct {
	Id         uint      `json:"id"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	FullName   string    `json:"fullName"`
	Avatar     string    `json:"avatar"`
	BirthDate  time.Time `json:"birthDate"`
	University string    `json:"university"`
	UniCode    string    `json:"university_code"`
	Npm        string    `json:"npm"`
	IsVerified bool      `json:"is_verified"`
}

func getUserInformation(userID uint) (UserInfoDTO, error) {
	db := database.GetDB()
	var user entity.User
	err := db.Preload("University").First(&user, "id = ?", userID).Error
	if err != nil {
		return UserInfoDTO{}, err
	}
	return UserInfoDTO{
		Id:         user.ID,
		Email:      user.Email,
		Username:   user.Username,
		FullName:   user.FullName,
		Avatar:     user.Avatar,
		BirthDate:  user.BirthDate,
		University: user.University.FullName,
		Npm:        user.Npm,
		UniCode:    user.University.UniCode,
		IsVerified: user.IsVerified,
	}, nil
}

func updateUserInformation(newUserData *UpdateUserDTO) error {
	db := database.GetDB()
	return db.Model(entity.User{}).Where("id = ?", newUserData.ID).Updates(entity.User{
		Email:        newUserData.Email,
		Username:     newUserData.Username,
		FullName:     newUserData.FullName,
		Avatar:       newUserData.Avatar,
		BirthDate:    newUserData.BirthDate.Time,
		UniversityID: newUserData.UniversityID,
		Npm:          newUserData.Npm,
	}).Error
}

func changeUserPassword(uid uint, pwdDTO *ChangePasswordDTO) error {
	db := database.GetDB()
	hashed, err := auth.HashPassword(pwdDTO.NewPassword)
	if err != nil {
		return err
	}
	return db.Model(entity.User{}).Where("id = ?", uid).Updates(entity.User{Password: hashed}).Error
}
