package main

import (
	"github.com/khilmi-aminudin/dvdrentalv1/router"
)

func main() {
	// r := gin.Default()
	// router.InitRouter(r, router.UserRouter, router.ActorRouter, router.AuthRouter)
	// r.Run()

	router.ServeRouter()
}
