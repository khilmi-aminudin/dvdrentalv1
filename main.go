package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khilmi-aminudin/dvdrentalv1/router"
)

func main() {
	r := gin.Default()
	router.UserRouter(r)
	router.ActorRouter(r)
	router.AuthRouter(r)
	r.Run()
}
