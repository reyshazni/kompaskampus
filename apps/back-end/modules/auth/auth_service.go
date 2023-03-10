package auth

import (
	"FindMyDosen/model/dto"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

func handleLogin(c echo.Context) error {
	user := new(dto.LoginUserDTO)
	if err := c.Bind(user); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.BaseDTO{
				Status:  http.StatusInternalServerError,
				Message: "Cannot decode Request Body",
				Data:    err.Error(),
			},
		)
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		fmt.Println(err)
		return c.JSON(
			http.StatusBadRequest,
			dto.BaseDTO{
				Status:  http.StatusBadRequest,
				Message: "Missing Field",
				Data:    err.Error(),
			},
		)
	}
	result, err := performUserLogin(user)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.BaseDTO{
				Status:  http.StatusInternalServerError,
				Message: "Fail to sign in user",
				Data:    err.Error(),
			},
		)
	}
	return c.JSON(http.StatusOK, dto.BaseDTO{
		Message: "User successfully login",
		Status:  http.StatusOK,
		Data:    result,
	})
}

func handleRegister(c echo.Context) error {
	newUser := new(dto.NewUserDTO)
	if err := c.Bind(newUser); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.BaseDTO{
				Status:  http.StatusInternalServerError,
				Message: "Cannot decode Request Body",
				Data:    err.Error(),
			},
		)
	}
	validate := validator.New()
	if err := validate.Struct(newUser); err != nil {
		fmt.Println(err)
		return c.JSON(
			http.StatusBadRequest,
			dto.BaseDTO{
				Status:  http.StatusBadRequest,
				Message: "Missing Field",
				Data:    err.Error(),
			},
		)
	}
	if newUser.Password != newUser.RePassword {
		return c.JSON(
			http.StatusBadRequest,
			dto.BaseDTO{
				Status:  http.StatusBadRequest,
				Message: "Password does not match",
			},
		)
	}
	err, result := performUserRegistration(newUser)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.BaseDTO{
				Status:  http.StatusInternalServerError,
				Message: "Fail to register user",
				Data:    err.Error(),
			},
		)
	}
	return c.JSON(
		http.StatusAccepted,
		dto.BaseDTO{
			Status:  http.StatusAccepted,
			Message: "User has been registered",
			Data:    result,
		},
	)
}
