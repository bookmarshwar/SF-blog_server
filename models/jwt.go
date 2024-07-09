package models

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	ID int `json:"ID"` //用户id
	jwt.RegisteredClaims
}
