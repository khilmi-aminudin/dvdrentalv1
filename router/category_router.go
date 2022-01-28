package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v4"
	"github.com/khilmi-aminudin/dvdrentalv1/controller"
	"github.com/khilmi-aminudin/dvdrentalv1/db"
	"github.com/khilmi-aminudin/dvdrentalv1/repository"
	"github.com/khilmi-aminudin/dvdrentalv1/service"
)

func CategoryRouter(r *gin.Engine) {
	var (
		db         *pgx.Conn                     = db.ConnectDBWithPGX()
		validator  *validator.Validate           = validator.New()
		repository repository.CategoryRepository = repository.NewcategoryRespository()
		service    service.CategoryService       = service.NewCategoryService(repository, db, validator)
		controller controller.CategoryController = controller.NewCategoryController(service)
	)

	r.POST("/api/category", controller.Create)
	r.PUT("/api/category/:id", controller.Update)
	r.DELETE("/api/category/:id", controller.Delete)

	r.GET("/api/category/:id", controller.FindById)
	r.GET("/api/categories", controller.FindAll)
}
