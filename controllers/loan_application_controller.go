package controllers

import (
	"fmt"

	"com.loan.demo/daos"
	"com.loan.demo/helpers"
	"com.loan.demo/services"
	"com.loan.demo/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ILoanApplicationController interface {
	IntiateApplication(*gin.Context)
	FetchBalanceSheet(*gin.Context)
	Submit(*gin.Context)
}

type LoanApplicationController struct {
	balanceSheetService   services.IBalanceSheetService
	decisionEngineService services.IDecisionEngineService
	applicationService    services.IApplicationService
}

func (lc LoanApplicationController) IntiateApplication(c *gin.Context) {
	var request daos.InitiateApplicationRequest
	c.BindJSON(&request)
	fmt.Println("Name", request.Name, "Year", request.YearEstablished, "Provider", request.Provider)
	id := lc.applicationService.CreateApplication(request)

	c.JSON(200, utils.JsonifyStruct(daos.InitiateResponseDao{
		ApplicationId: id,
	}))
}

func (lc LoanApplicationController) FetchBalanceSheet(c *gin.Context) {
	var request daos.FetchBalanceSheetRequest
	c.BindQuery(&request)
	fmt.Println("REQUEST", request.Name, request.Year)
	c.JSON(200, utils.ArrayResponse(lc.balanceSheetService.GetBalanceSheet(request)))
}

func (lc LoanApplicationController) Submit(c *gin.Context) {
	var request daos.SubmitApplicationRequest
	c.BindJSON(&request)
	applicationSummary := lc.applicationService.SummariseApplication(request.ApplicationId, request.ProfitLossInLastYear, request.AverageAssetValue)
	decisionEngineOutput := helpers.ApplyRules(applicationSummary, request.Name, request.YearEstablished, request.LoanAmount)
	status, message := lc.decisionEngineService.CheckEligibility(decisionEngineOutput)
	c.JSON(200, utils.JsonifyStruct(daos.SubmissionDao{
		Status:  status,
		Message: message,
	}))
}

func NewLoanApplicationController(balanceSheetUrl string, decisionEngineUrl string, db *gorm.DB) ILoanApplicationController {
	return LoanApplicationController{
		balanceSheetService:   services.NewBalanceSheetService(balanceSheetUrl),
		decisionEngineService: services.NewDecsionEngineService(decisionEngineUrl),
		applicationService:    services.NewApplicationService(db),
	}
}
