package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"github.com/khilmi-aminudin/dvdrentalv1/models/entity"
	"github.com/khilmi-aminudin/dvdrentalv1/models/web"
	"github.com/khilmi-aminudin/dvdrentalv1/service"
)

type UserController interface {
	Register(c *gin.Context)
	FindAll(c *gin.Context)
	Login(c *gin.Context)
}

type userController struct {
	Service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		Service: service,
	}
}

func (controller *userController) Register(c *gin.Context) {
	var request web.RequestCreateUser
	err := c.BindJSON(&request)
	helper.PanicIfError(err)

	response := controller.Service.Create(c.Request.Context(), request)
	c.JSON(http.StatusOK, response)

}

func (controller *userController) FindAll(c *gin.Context) {
	response := controller.Service.FindAll(c.Request.Context())
	c.JSON(http.StatusOK, response)
}

func (controller *userController) Login(c *gin.Context) {
	var credential web.LoginCredential
	err := c.BindJSON(&credential)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	response := controller.Service.FindByUsername(c.Request.Context(), credential.Username)

	if response.Data == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": fmt.Sprintf("user with username %s not found", credential.Username),
		})
		return
	}

	credential.Passowrd = helper.NewSHA256([]byte(credential.Passowrd))

	user, ok := response.Data.(entity.Users)

	if !ok {

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": "-------",
		})
		return
	}

	if user.Username == credential.Username && user.Passowrd == credential.Passowrd {
		c.JSON(http.StatusOK, gin.H{
			"code":     http.StatusOK,
			"status":   "Authorized",
			"message":  "login succes",
			"is_login": true,
		})
	}

}
