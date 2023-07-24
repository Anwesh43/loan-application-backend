package main

import (
	"com.loan.demo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routeHandler := routes.NewRouteHandler()
	routeHandler.PopulateRouteMap()
	routeHandler.InitiateRoutes(r)
	r.Run()
}
