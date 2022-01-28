package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/khilmi-aminudin/dvdrentalv1/helper"
)

func InitRouter(r *gin.Engine, router ...func(*gin.Engine)) {
	for _, route := range router {
		route(r)
	}
}

func ServeRouter() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	r := gin.Default()
	InitRouter(r, UserRouter, ActorRouter, AuthRouter, CategoryRouter)
	r.Run()
}
