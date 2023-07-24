package services

type IDecisionEngineService interface {
	CheckEligibility(details interface{}) bool
}

type DecisionEngineService struct {
	url string
}

func (d DecisionEngineService) CheckEligibility(details interface{}) bool {
	return true
}

func NewDecsionEngineService(url string) DecisionEngineService {
	return DecisionEngineService{
		url: url,
	}
}
