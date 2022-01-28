package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"github.com/khilmi-aminudin/dvdrentalv1/models/web"
	"github.com/khilmi-aminudin/dvdrentalv1/service"
)

type CategoryController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
}

type categoryController struct {
	service service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &categoryController{
		service: service,
	}
}

func (controller *categoryController) Create(c *gin.Context) {
	var request web.RequestCreateCategory

	err := c.BindJSON(&request)

	helper.PanicIfError(err)

	response := controller.service.Create(c.Request.Context(), request)

	c.JSON(http.StatusOK, response)
}

func (controller *categoryController) Update(c *gin.Context) {
	var request web.RequestUpdateCategory

	err := c.BindJSON(&request)

	helper.PanicIfError(err)

	response := controller.service.Update(c.Request.Context(), request)

	c.JSON(http.StatusOK, response)

}

func (controller *categoryController) Delete(c *gin.Context) {
	id := c.Param("id")

	categoryId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	response := controller.service.Delete(c.Request.Context(), int64(categoryId))

	c.JSON(http.StatusOK, response)
}

func (controller *categoryController) FindById(c *gin.Context) {
	id := c.Param("id")

	categoryId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	response := controller.service.FindById(c.Request.Context(), int64(categoryId))

	c.JSON(http.StatusOK, response)
}

func (controller *categoryController) FindAll(c *gin.Context) {
	response := controller.service.FindAll(c.Request.Context())

	c.JSON(http.StatusOK, response)
}
