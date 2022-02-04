package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v4"
	"github.com/khilmi-aminudin/dvdrentalv1/controller"
	"github.com/khilmi-aminudin/dvdrentalv1/db"
	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"github.com/khilmi-aminudin/dvdrentalv1/models/entity"
	"github.com/khilmi-aminudin/dvdrentalv1/repository"
	"github.com/khilmi-aminudin/dvdrentalv1/service"
)

func ExcelRouter(r *gin.Engine) {
	var (
		db              *pgx.Conn                  = db.ConnectDBWithPGX()
		validator       *validator.Validate        = validator.New()
		actorRepository repository.ActorRepository = repository.NewActorRepository()
		actorService    service.ActorService       = service.NewActorService(actorRepository, db, validator)
		ctx             context.Context            = context.Background()
	)

	response := actorService.FindAll(ctx)

	var actors, ok = response.Data.([]entity.Actor)
	if !ok {
		helper.LoggerInit().Warn("Cannt convert response actors")
	}

	var (
		excelService service.ExcelService       = service.NewExcelService(actors)
		controller   controller.ExcelController = controller.NewExcelController(excelService)
	)

	r.GET("/api/excel/actors", controller.GenerateExcel)
}
