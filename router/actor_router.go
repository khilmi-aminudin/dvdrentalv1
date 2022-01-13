package router

import (
	"github.com/khilmi-aminudin/dvdrentalv1/controller"
	"github.com/khilmi-aminudin/dvdrentalv1/db"
	"github.com/khilmi-aminudin/dvdrentalv1/repository"
	"github.com/khilmi-aminudin/dvdrentalv1/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v4"
)

func ActorRouter(r *gin.Engine) {
	var (
		db         *pgx.Conn                  = db.ConnectDBWithPGX()
		validator  *validator.Validate        = validator.New()
		repository repository.ActorRepository = repository.NewActorRepository()
		service    service.ActorService       = service.NewActorService(repository, db, validator)
		controller controller.ActorController = controller.NewActorController(service)
	)

	r.GET("/api/actor/:id", controller.FindById)
	r.GET("/api/actor", controller.FindAll)
	r.POST("/api/actor", controller.Create)
	r.DELETE("/api/actor/:id", controller.Delete)
	r.PUT("/api/actor", controller.Update)
	r.GET("/api/actor/search", controller.Search)

}
