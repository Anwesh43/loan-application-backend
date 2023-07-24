package services

import "com.loan.demo/daos"

type IDecisionEngineService interface {
	CheckEligibility(daos.DecisionEngineOutput) (bool, string)
}

type DecisionEngineService struct {
	url string
}

func (d DecisionEngineService) CheckEligibility(details daos.DecisionEngineOutput) (bool, string) {
	if details.PreAssessment == "60%" || details.PreAssessment == "100%" {
		return true, "Loan Application Accepted"
	}
	return false, "Loan Application Rejected"
}

func NewDecsionEngineService(url string) DecisionEngineService {
	return DecisionEngineService{
		url: url,
	}
}
