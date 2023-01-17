package user

import (
	"FindMyDosen/model/dto"
	"FindMyDosen/model/entity"
	"FindMyDosen/modules/auth"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetUserInfoDTO struct {
	userID uint `json:"user_id"`
}

type UpdateUserDTO struct {
	ID           uint          `json:"id" validate:"required"`
	Email        string        `json:"email" validate:"required,email"`
	Username     string        `json:"username" validate:"required"`
	FullName     string        `json:"fullName" validate:"required"`
	Avatar       string        `json:"avatar" validate:"omitempty,required"`
	BirthDate    dto.BirthDate `json:"birth_date" validate:"required"`
	UniversityID int           `json:"university_id" validate:"required"`
	Npm          string        `json:"npm" validate:"required"`
}

type ChangePasswordDTO struct {
	OldPassword   string `json:"old_password" validate:"required"`
	ReOldPassword string `json:"re_old_password" validate:"required"`
	NewPassword   string `json:"new_password" validate:"required"`
	ReNewPassword string `json:"re_new_password" validate:"required"`
}

func handleChangePassword(c echo.Context) error {
	pwdDTO := new(ChangePasswordDTO)
	if err := c.Bind(&pwdDTO); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BaseDTO{
			Message: "Cannot Bind JSON",
			Status:  http.StatusInternalServerError,
		})
	}
	validate := validator.New()
	if err := validate.Struct(pwdDTO); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dto.BaseDTO{
				Status:  http.StatusBadRequest,
				Message: "Missing Field",
				Data:    err.Error(),
			},
		)
	}
	if pwdDTO.NewPassword != pwdDTO.ReNewPassword || pwdDTO.ReOldPassword != pwdDTO.OldPassword {
		return c.JSON(
			http.StatusBadRequest,
			dto.BaseDTO{
				Status:  http.StatusBadRequest,
				Message: "Password does not match!",
			},
		)
	}
	
	if msg, state := auth.IsValidPassword(pwdDTO.NewPassword); !state {
		return c.JSON(
			http.StatusBadRequest,
			dto.BaseDTO{
				Status:  http.StatusBadRequest,
				Message: msg,
			},
		)
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*entity.JwtClaims)
	if err := changeUserPassword(claims.Uid, pwdDTO); err != nil {
		return c.JSON(
			http.StatusNotFound,
			dto.BaseDTO{
				Message: "Cannot found user",
				Status:  http.StatusNotFound,
			},
		)
	}
	return c.JSON(http.StatusOK, dto.BaseDTO{
		Message: "User password has been updated",
		Status:  http.StatusOK,
	})
}

func handleGetUserInformation(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*entity.JwtClaims)
	locDto := new(GetUserInfoDTO)
	if err := c.Bind(&locDto); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BaseDTO{
			Message: "Cannot Bind JSON",
			Status:  http.StatusInternalServerError,
		})
	}
	result, err := getUserInformation(claims.Uid)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.BaseDTO{
			Message: "Cannot find user",
			Status:  404,
		})
	}
	return c.JSON(http.StatusOK, dto.BaseDTO{
		Message: "Ok",
		Status:  http.StatusOK,
		Data:    result,
	})
}

func handleUpdateUserInformation(c echo.Context) error {
	updateUser := new(UpdateUserDTO)
	if err := c.Bind(&updateUser); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BaseDTO{
			Message: "Cannot Bind JSON",
			Status:  http.StatusInternalServerError,
		})
	}
	validate := validator.New()
	if err := validate.Struct(updateUser); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dto.BaseDTO{
				Status:  http.StatusBadRequest,
				Message: "Missing Field",
				Data:    err.Error(),
			},
		)
	}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*entity.JwtClaims)
	if claims.Uid != updateUser.ID {
		println(claims.Uid, updateUser.ID)
		return c.JSON(
			http.StatusUnauthorized,
			dto.BaseDTO{
				Status:  http.StatusUnauthorized,
				Message: "User ID not Valid!",
			},
		)
	}
	if err := updateUserInformation(updateUser); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BaseDTO{
			Status:  http.StatusInternalServerError,
			Message: "Cannot update user information",
			Data:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.BaseDTO{
		Message: "User information has been updated",
		Status:  http.StatusOK,
		Data:    updateUser,
	})
}
