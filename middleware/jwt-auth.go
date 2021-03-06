package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/fajarbc/learn-gin/service"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := validateToken(ctx)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[ID]: ", claims["id"])
			// log.Println("Claims[Name]: ", claims["name"])
			// log.Println("Claims[Admin]: ", claims["admin"])
			// log.Println("Claims[Issuer]: ", claims["iss"])
			// log.Println("Claims[IssuedAt]: ", claims["iat"])
			// log.Println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusBadRequest)
		}
	}
}

func GetTokenClaim(ctx *gin.Context) (jwt.MapClaims, error) {
	token, err := validateToken(ctx)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		return claims, nil
	} else {
		log.Println("Failed to GetTokenClaim")
		log.Println(err)
		return nil, err
	}
}

func validateToken(ctx *gin.Context) (*jwt.Token, error) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	// log.Printf("authHeader: %s\n", authHeader)
	if authHeader == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]

	token, err := service.NewJWTService().ValidateToken(tokenString)
	return token, err
}
