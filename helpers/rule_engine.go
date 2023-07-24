package helpers

import (
	"com.loan.demo/daos"
	"com.loan.demo/models"
)

func ApplyRules(applicationSummary models.ApplicationSummary, name string, year string, loanAmount int) daos.DecisionEngineOutput {
	preAssessment := "0%"
	if applicationSummary.PLInLastYear > 0 {
		preAssessment = "60%"
	}
	if applicationSummary.AverageAssetValue > loanAmount {
		preAssessment = "100%"
	}
	return daos.DecisionEngineOutput{
		PreAssessment:        preAssessment,
		Name:                 name,
		YearEstablished:      year,
		ApplicationId:        applicationSummary.CurrApplication.Id,
		ProfitLossInLastYear: applicationSummary.PLInLastYear,
	}
}
