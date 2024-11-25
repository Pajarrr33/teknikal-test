package controller

import (
	"fmt"
	"net/http"
	"strings"
	"teknikal-test/delivery/middleware"
	"teknikal-test/entity/request"
	"teknikal-test/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUc usecase.AuthUsecase
	rg     *gin.RouterGroup
	authMd middleware.AuthMiddleware
}

func (a *AuthController) Login(ctx *gin.Context) {
	var loginRequest request.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponse, err := a.authUc.Login(loginRequest)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if loginResponse.AccessToken == ""  {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, loginResponse)
}

func (a *AuthController) Register(ctx *gin.Context) {
	var registerRequest request.RegisterRequest
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := a.authUc.Register(registerRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, customer)
}

func (a *AuthController) Logout(ctx *gin.Context) {
	var authHeader middleware.AuthHeader

	if err := ctx.ShouldBindHeader(&authHeader); err != nil {
		fmt.Println("Error while binding token in logout", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := strings.TrimPrefix(authHeader.AuthorizationHeader, "Bearer ")
	if token == "" {
		fmt.Println("Error while trim string in logout")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := ctx.GetString("id")

	err := a.authUc.Logout(token,id)
	if err != nil {
		fmt.Println("Error while logout in usecase", err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (a *AuthController) Route() {
	a.rg.POST("/login", a.Login)
	a.rg.POST("/register", a.Register)
	a.rg.POST("/logout", a.authMd.RequireToken(), a.Logout)
}

func NewAuthController(authUc usecase.AuthUsecase, rg *gin.RouterGroup, authMd middleware.AuthMiddleware) *AuthController {
	return &AuthController{authUc, rg, authMd}
}
