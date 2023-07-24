package controllers

import (
	"com.loan.demo/daos"
	"com.loan.demo/services"
	"github.com/gin-gonic/gin"
)

type ILoanApplicationController interface {
	IntiateApplication(*gin.Context)
	FetchBalanceSheet(*gin.Context)
	Submit(*gin.Context)
}

type LoanApplicationController struct {
	balanceSheetService   services.IBalanceSheetService
	decisionEngineService services.IDecisionEngineService
}

func (lc LoanApplicationController) IntiateApplication(c *gin.Context) {
	c.JSON(200, daos.InitiateResponseDao{
		ApplicationId: "application1",
	})
}

func (lc LoanApplicationController) FetchBalanceSheet(c *gin.Context) {
	c.JSON(200, daos.BalanceSheetResponseDao{
		Year:          "2020",
		Month:         10,
		ProfitsOrLoss: 1150,
		AssetsValue:   1234,
	})
}

func (lc LoanApplicationController) Submit(c *gin.Context) {
	c.JSON(200, daos.SubmissionDao{
		Status:  true,
		Message: "Submitted Successfully",
	})
}

func NewLoanApplicationController(balanceSheetUrl string, decisionEngineUrl string) ILoanApplicationController {
	return LoanApplicationController{
		balanceSheetService:   services.NewBalanceSheetService(balanceSheetUrl),
		decisionEngineService: services.NewDecsionEngineService(decisionEngineUrl),
	}
}
