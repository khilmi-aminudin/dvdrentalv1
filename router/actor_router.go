package router

import (
	"github.com/khilmi-aminudin/dvdrentalv1/controller"
	"github.com/khilmi-aminudin/dvdrentalv1/db"
	"github.com/khilmi-aminudin/dvdrentalv1/middlewares"
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

	authorized := r.Group("/api/actor", middlewares.JWTAuthentication())

	authorized.GET("/:id", controller.FindById)
	authorized.GET("/", controller.FindAll)
	authorized.POST("/", controller.Create)
	authorized.DELETE("/:id", controller.Delete)
	authorized.PUT("/", controller.Update)
	authorized.GET("/search", controller.Search)

}
