package middlewares

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/khilmi-aminudin/dvdrentalv1/service"
)

func JWTAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHENA = "bearersch "
		authenticationheader := c.GetHeader("Authentication")
		tokenString := authenticationheader[len(BEARER_SCHENA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]\t: ", claims["name"])
			log.Println("Claims[IsSignin]\t: ", claims["is_signin"])
			log.Println("Claims[Issuer]\t: ", claims["iss"])
			log.Println("Claims[IssuedAt]\t: ", claims["iat"])
			log.Println("Claims[ExpiresAt]\t: ", claims["exp"])
		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
