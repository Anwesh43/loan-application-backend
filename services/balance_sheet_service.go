package services

import (
	"com.loan.demo/daos"
	"com.loan.demo/helpers"
)

type IBalanceSheetService interface {
	GetBalanceSheet(daos.FetchBalanceSheetRequest) []daos.BalanceSheetResponseDao
}

type BalanceSheetService struct {
	url string
}

func (b BalanceSheetService) GetBalanceSheet(request daos.FetchBalanceSheetRequest) []daos.BalanceSheetResponseDao {
	//TODO: use url to get data from an accounting software endpooint
	responses := make([]daos.BalanceSheetResponseDao, 0)
	return helpers.GetMockBalanceSheetData(request.Name, request.Year, responses)
}

func NewBalanceSheetService(url string) IBalanceSheetService {
	return BalanceSheetService{
		url: url,
	}
}
