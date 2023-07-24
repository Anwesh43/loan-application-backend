package helpers

import "com.loan.demo/daos"

func GetMockBalanceSheetData(name string, year string, responses []daos.BalanceSheetResponseDao) []daos.BalanceSheetResponseDao {
	responses = append(responses, daos.BalanceSheetResponseDao{
		Year:          year,
		Month:         12,
		ProfitsOrLoss: 250000,
		AssetsValue:   1234,
	})
	responses = append(responses, daos.BalanceSheetResponseDao{
		Year:          year,
		Month:         11,
		ProfitsOrLoss: 1150,
		AssetsValue:   5789,
	})
	responses = append(responses, daos.BalanceSheetResponseDao{
		Year:          year,
		Month:         10,
		ProfitsOrLoss: 2500,
		AssetsValue:   22345,
	})
	responses = append(responses, daos.BalanceSheetResponseDao{
		Year:          year,
		Month:         9,
		ProfitsOrLoss: -187000,
		AssetsValue:   223452,
	})
	return responses
}
