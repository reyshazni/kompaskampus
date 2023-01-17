package entity

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	Uid        uint
	IsVerified bool
	jwt.RegisteredClaims
}
