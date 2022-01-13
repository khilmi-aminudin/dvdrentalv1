package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(authHeader):]

		fmt.Println(tokenString)
	}
}