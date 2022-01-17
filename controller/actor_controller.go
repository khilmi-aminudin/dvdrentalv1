package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"github.com/khilmi-aminudin/dvdrentalv1/models/web"
	"github.com/khilmi-aminudin/dvdrentalv1/service"
)

type ActorController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Search(c *gin.Context)
}

type actorController struct {
	Service service.ActorService
}

func NewActorController(service service.ActorService) ActorController {
	return &actorController{
		Service: service,
	}
}

func (controller *actorController) Create(c *gin.Context) {
	var requestCreateActor web.RequestCreateActor
	err := c.BindJSON(&requestCreateActor)
	helper.PanicIfError(err)

	response := controller.Service.Create(c.Request.Context(), requestCreateActor)
	c.JSON(http.StatusOK, response)

}

func (controller *actorController) Update(c *gin.Context) {
	var requestUpdateActor web.RequestUpdateActor
	err := c.BindJSON(&requestUpdateActor)
	helper.PanicIfError(err)

	response := controller.Service.Update(c.Request.Context(), requestUpdateActor)
	c.JSON(http.StatusOK, response)
}

func (controller *actorController) Delete(c *gin.Context) {
	actorId := c.Param("id")
	id, err := strconv.Atoi(actorId)
	helper.PanicIfError(err)

	response := controller.Service.Delete(c.Request.Context(), int64(id))

	if response.Data == nil {
		c.JSON(http.StatusNotFound, web.ResponseWeb{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data: map[string]string{
				"message": fmt.Sprintf("data with id %d not exist", id),
			},
		})
		return
	}

	c.JSON(http.StatusOK, web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data: map[string]string{
			"message": fmt.Sprintf("data with id %d was deleted", id),
		},
	})
}

func (controller *actorController) FindById(c *gin.Context) {
	actorId := c.Param("id")
	id, err := strconv.Atoi(actorId)
	helper.PanicIfError(err)

	response := controller.Service.FindById(c.Request.Context(), int64(id))
	if fmt.Sprintf("%T", response.Data) == "string" {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (controller *actorController) FindAll(c *gin.Context) {
	response := controller.Service.FindAll(c.Request.Context())
	c.JSON(http.StatusOK, response)
}

func (controller *actorController) Search(c *gin.Context) {
	key := c.Query("key")
	response := controller.Service.Search(c.Request.Context(), key)
	c.JSON(http.StatusOK, response)
}
