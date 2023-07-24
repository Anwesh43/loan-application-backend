package daos

type InitiateApplicationRequest struct {
	Name            string `json:"name"`
	YearEstablished string `json:"year_established"`
	Provider        string `json:"provider"`
}

type FetchBalanceSheetRequest struct {
	Name string `form:"name"`
	Year string `form:"year"`
}

type SubmitApplicationRequest struct {
	Name                 string `json:"name"`
	YearEstablished      string `json:"year_established"`
	ProfitLossInLastYear int    `json:"plInLastYear"`
	AverageAssetValue    int    `json:"averageAssetValue"`
	ApplicationId        int    `json:"applicationId"`
	LoanAmount           int    `json:"loanAmount"`
}

type DecisionEngineOutput struct {
	Name                 string
	YearEstablished      string
	ProfitLossInLastYear int
	ApplicationId        int
	PreAssessment        string
}
