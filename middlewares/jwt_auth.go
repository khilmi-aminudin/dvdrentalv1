package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"github.com/khilmi-aminudin/dvdrentalv1/service"
	"github.com/sirupsen/logrus"
)

func JWTAuthentication() gin.HandlerFunc {
	logger := helper.LoggerInit()
	return func(c *gin.Context) {
		const BEARER_SCHENA = "bearersch "
		authenticationheader := c.GetHeader("Authentication")
		tokenString := authenticationheader[len(BEARER_SCHENA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			logger.WithFields(logrus.Fields{
				"Claims[Name]":      claims["name"],
				"Claims[IsSignin]":  claims["is_signin"],
				"Claims[Issuer]":    claims["iss"],
				"Claims[IssuedAt]":  claims["iat"],
				"Claims[ExpiresAt]": claims["exp"],
			}).Info(claims)
		} else {
			logger.WithField("error", err).Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
