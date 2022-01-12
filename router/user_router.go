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

func UserRouter(r *gin.Engine) {
	var (
		db         *pgx.Conn                 = db.ConnectDBWithPGX()
		validator  *validator.Validate       = validator.New()
		repository repository.UserRepository = repository.NewuserRepository()
		service    service.UserService       = service.NewuserService(repository, db, validator)
		controller controller.UserController = controller.NewUserController(service)
	)

	r.GET("/api/users", controller.FindAll)
	r.POST("/api/user", controller.Register)
	r.POST("/api/user/login", controller.Login)
}
