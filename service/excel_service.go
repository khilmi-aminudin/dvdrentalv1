package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"github.com/khilmi-aminudin/dvdrentalv1/models/entity"
	"github.com/khilmi-aminudin/dvdrentalv1/models/web"
)

type ExcelService interface {
	GenerateExcel() web.ResponseWeb
}

type excelService struct {
	actors []entity.Actor
}

func NewExcelService(actors []entity.Actor) ExcelService {
	return &excelService{
		actors: actors,
	}
}

func (service *excelService) GenerateExcel() web.ResponseWeb {
	xlsx := excelize.NewFile()

	sheet1Name := "Sheet One"

	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)
	xlsx.SetCellValue(sheet1Name, "A1", "Actor Id")
	xlsx.SetCellValue(sheet1Name, "B1", "First Name")
	xlsx.SetCellValue(sheet1Name, "C1", "Last Name")
	xlsx.SetCellValue(sheet1Name, "D1", "Last Update")

	err := xlsx.AutoFilter(sheet1Name, "A1", "C1", "")
	if err != nil {
		helper.LoggerInit().Warn("ERROR", err.Error())
		log.Fatal("ERROR", err.Error())
		return web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   fmt.Sprintf("Error : %s", err.Error()),
		}
	}

	for i, actor := range service.actors {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), actor.ActorId)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), actor.FirstName)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), actor.LastName)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+2), actor.LastUpdate)
	}

	err = xlsx.SaveAs("./file1.xlsx")
	if err != nil {
		helper.LoggerInit().Warn("ERROR", err.Error())
		log.Fatal("ERROR", err.Error())

		return web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   fmt.Sprintf("Error : %s", err.Error()),
		}
	}
	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
	}

}
