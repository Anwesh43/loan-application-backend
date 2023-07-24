package main

import (
	"log"

	"com.loan.demo/helpers"
	"com.loan.demo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	helpers.LoadEnv()
	db, err := helpers.ConnectToDb()
	routeHandler := routes.NewRouteHandler(db)
	routeHandler.PopulateRouteMap()
	routeHandler.InitiateRoutes(r)

	if err == nil {
		log.Println("Connected to database")
	}
	r.Run()
}
