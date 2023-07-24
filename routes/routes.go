package routes

import (
	"com.loan.demo/controllers"
	"github.com/gin-gonic/gin"
)

type IRouteHandler interface {
	InitiateRoutes(g *gin.Engine)
	PopulateRouteMap()
}

type Route struct {
	handler gin.HandlerFunc
	method  string
}

type LoanRouteHandler struct {
	controller controllers.ILoanApplicationController
	routeMaps  map[string]Route
}

func (r LoanRouteHandler) PopulateRouteMap() {
	r.routeMaps["/initiate"] = Route{
		handler: r.controller.IntiateApplication,
		method:  "POST",
	}
	r.routeMaps["/fetch-balance-sheet"] = Route{
		handler: r.controller.FetchBalanceSheet,
		method:  "GET",
	}
	r.routeMaps["/submit"] = Route{
		handler: r.controller.Submit,
		method:  "POST",
	}
}

func (r LoanRouteHandler) InitiateRoutes(g *gin.Engine) {
	for k, v := range r.routeMaps {
		if v.method == "GET" {
			g.GET(k, v.handler)
		}
		if v.method == "POST" {
			g.POST(k, v.handler)
		}
	}
}

func NewRouteHandler() IRouteHandler {
	return LoanRouteHandler{
		routeMaps:  make(map[string]Route),
		controller: controllers.NewLoanApplicationController("", ""),
	}
}
