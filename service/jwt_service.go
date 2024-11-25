package service

import (
	"fmt"
	"teknikal-test/config"
	"teknikal-test/entity"
	"teknikal-test/entity/response"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(customer entity.Customer) (response.LoginResponse, error)
	ValidateToken(token string) (jwt.MapClaims, error)
}

type jwtService struct {
	config config.JwtConfig
}

func (j *jwtService) GenerateToken(customer entity.Customer) (response.LoginResponse, error) {
	accessClaims := entity.MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.Expire)),
			Issuer:    j.config.Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Id:    customer.Id,
		Email: customer.Email,
	}

	accessToken := jwt.NewWithClaims(j.config.Method, accessClaims)
	accessTokenString, err := accessToken.SignedString(j.config.SecretKey)
	if err != nil {
		return response.LoginResponse{}, fmt.Errorf("failed to generate access token: %v", err)
	}


	return response.LoginResponse{AccessToken: accessTokenString}, nil
}

func (j *jwtService) ValidateToken(token string) (jwt.MapClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &entity.MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tokenClaims.Claims.(*entity.MyCustomClaims)
	if !ok || !tokenClaims.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return jwt.MapClaims{
		"id":    claims.Id,
		"email": claims.Email,
	},nil
}

func NewJWTService(config config.JwtConfig) JWTService {
	return &jwtService{config: config}
}