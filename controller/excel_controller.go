package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khilmi-aminudin/dvdrentalv1/service"
)

type ExcelController interface {
	GenerateExcel(c *gin.Context)
}

type excelController struct {
	service service.ExcelService
}

func NewExcelController(excelService service.ExcelService) ExcelController {
	return &excelController{
		service: excelService,
	}
}

func (controller *excelController) GenerateExcel(c *gin.Context) {
	response := controller.service.GenerateExcel()

	if response.Data != nil {
		c.JSON(http.StatusBadRequest, response)
	}

	c.JSON(http.StatusOK, response)
}
