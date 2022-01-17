package router

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v4"
	"github.com/khilmi-aminudin/dvdrentalv1/controller"
	"github.com/khilmi-aminudin/dvdrentalv1/db"
	"github.com/khilmi-aminudin/dvdrentalv1/models/entity"
	"github.com/khilmi-aminudin/dvdrentalv1/models/web"
	"github.com/khilmi-aminudin/dvdrentalv1/repository"
	"github.com/khilmi-aminudin/dvdrentalv1/service"
)

func AuthRouter(r *gin.Engine) {
	var (
		db          *pgx.Conn                 = db.ConnectDBWithPGX()
		validator   *validator.Validate       = validator.New()
		repository  repository.UserRepository = repository.NewuserRepository()
		userservice service.UserService       = service.NewuserService(repository, db, validator)
		ctx         context.Context           = context.Background()
		response    web.ResponseWeb           = userservice.FindAll(ctx)
	)

	var users, ok = response.Data.([]entity.Users)
	if !ok {
		fmt.Println(users)
	}
	var service service.Authentication = service.NewAuthentication(users)
	var controller controller.AuthController = controller.NewAuthController(service)

	r.POST("/api/auth/login", controller.Login)
}
