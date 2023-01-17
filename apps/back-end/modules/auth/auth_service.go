package auth

import (
	"FindMyDosen/config"
	"FindMyDosen/model/dto"
	"FindMyDosen/model/entity"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"time"
)

type RefreshDTO struct {
	UserID       uint   `json:"user_id" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func handleRefresh(c echo.Context) error {
	refreshDTO := new(RefreshDTO)
	if err := c.Bind(refreshDTO); err != nil {
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
	if err := validate.Struct(refreshDTO); err != nil {
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
	result, err := performRefreshToken(refreshDTO.UserID, refreshDTO.RefreshToken)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.BaseDTO{
				Status:  http.StatusInternalServerError,
				Message: "Fail to generate Token",
				Data:    err.Error(),
			},
		)
	}
	return c.JSON(http.StatusOK, dto.BaseDTO{
		Message: "Generate new Token",
		Status:  http.StatusOK,
		Data:    result,
	})
}

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

	if msg, state := IsValidPassword(newUser.Password); !state {
		return c.JSON(
			http.StatusBadRequest,
			dto.BaseDTO{
				Status:  http.StatusBadRequest,
				Message: msg,
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

func IsValidPassword(password string) (string, bool) {
	if len(password) < 8 || len(password) > 15 {
		return "Password must be between 8 to 15 character long", false
	}
	var special = regexp.MustCompile(`[!@#\$%^&\*\(\)_+\-=\[\]{};':"\\|,.<>\/?]`)
	var uppercase = regexp.MustCompile(`[A-Z]`)
	var lowercase = regexp.MustCompile(`[a-z]`)
	var numeric = regexp.MustCompile(`[0-9]`)

	if !special.MatchString(password) {
		return "Password must contain at least one special character", false
	}
	if !uppercase.MatchString(password) {
		return "Password must contain at least one uppercase letter", false
	}
	if !lowercase.MatchString(password) {
		return "Password must contain at least one lowercase letter", false
	}
	if !numeric.MatchString(password) {
		return "Password must contain at least one number", false
	}

	return "", true
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPassword(password string, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err
}

func generateToken(uid uint, isVerified bool) (string, error) {
	claims := entity.JwtClaims{
		uid,
		isVerified,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetJWTSecret()))
}

func handleVerifyUser(c echo.Context) error {
	_ = c.Param("code")
	return c.String(http.StatusOK, "s")
}
