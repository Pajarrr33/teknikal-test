package entity

import "github.com/golang-jwt/jwt/v5"

type MyCustomClaims struct {
	jwt.RegisteredClaims
	Id string `json:"id"`
	Email string `json:"userId"`
}