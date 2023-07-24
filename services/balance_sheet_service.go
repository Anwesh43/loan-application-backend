package services

type IBalanceSheetService interface {
	GetBalanceSheet(details interface{}) string
}

type BalanceSheetService struct {
	url string
}

func (b BalanceSheetService) GetBalanceSheet(details interface{}) string {
	return ""
}

func NewBalanceSheetService(url string) IBalanceSheetService {
	return BalanceSheetService{
		url: url,
	}
}
