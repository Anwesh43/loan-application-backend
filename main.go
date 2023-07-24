package main

import (
	"log"

	"time"

	"com.loan.demo/helpers"
	"com.loan.demo/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
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
