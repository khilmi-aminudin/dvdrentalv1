package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"github.com/khilmi-aminudin/dvdrentalv1/models/dto"
	"github.com/khilmi-aminudin/dvdrentalv1/models/web"
	"github.com/khilmi-aminudin/dvdrentalv1/service"
)

type AuthController interface {
	Login(c *gin.Context)
}

type authcontroller struct {
	auth       service.Authentication
	jwtservice service.JWTService
}

func NewAuthController(authservice service.Authentication, jwtsrvc service.JWTService) AuthController {
	return &authcontroller{
		auth:       authservice,
		jwtservice: jwtsrvc,
	}
}

func (controller *authcontroller) Login(c *gin.Context) {
	var credential dto.Credential
	err := c.BindJSON(&credential)

	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data: gin.H{
				"message": err.Error(),
			},
		})
	}

	hashedpassword := helper.NewSHA256([]byte(credential.Password))
	isSignin := controller.auth.Login(credential.Username, hashedpassword)
	if !isSignin {

		c.JSON(http.StatusUnauthorized, web.ResponseWeb{
			Code:   http.StatusUnauthorized,
			Status: "Bad Request",
			Data:   credential,
		})
	} else {
		token := controller.jwtservice.GenerateToken(credential.Username, isSignin)
		c.JSON(http.StatusOK, web.ResponseWeb{
			Code:   http.StatusOK,
			Status: "Authorized",
			Data: gin.H{
				"credentials": credential,
				"token":       token,
			},
		})
	}
}
