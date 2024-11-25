package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"teknikal-test/repository"
	"teknikal-test/service"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	RequireToken() gin.HandlerFunc
}

type authMiddleware struct {
	jwtService service.JWTService
	expRepo 	repository.ExpiredRepository
}

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func (a *authMiddleware) RequireToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var authHeader AuthHeader
		if err := ctx.ShouldBindHeader(&authHeader); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		tokenHeader := strings.TrimPrefix(authHeader.AuthorizationHeader, "Bearer ")
		if tokenHeader == "" {
			fmt.Println(tokenHeader)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		_, err := a.expRepo.GetExpiredByToken(tokenHeader)
		if err != sql.ErrNoRows && err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		
		claims, err := a.jwtService.ValidateToken(tokenHeader)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("id", claims["id"])
		ctx.Set("email", claims["email"])

		if claims["email"] == "" || claims["id"] == "" {
			fmt.Println("claims is empty")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}


func NewAuthMiddleware(jwtService service.JWTService, expRepo repository.ExpiredRepository) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService, expRepo: expRepo}
}