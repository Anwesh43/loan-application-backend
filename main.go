package main

import (
	"log"

	"com.loan.demo/helpers"
	"com.loan.demo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routeHandler := routes.NewRouteHandler()
	routeHandler.PopulateRouteMap()
	routeHandler.InitiateRoutes(r)
	helpers.LoadEnv()
	_, err := helpers.ConnectToDb()
	if err == nil {
		log.Println("Connected to database")
	}
	r.Run()
}
