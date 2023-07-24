package daos

type InitiateResponseDao struct {
	ApplicationId int `json:"applicationId"`
}

type BalanceSheetResponseDao struct {
	Year          string `json:"year"`
	Month         int    `json:"month"`
	ProfitsOrLoss int    `json:"profitsOrLoss"`
	AssetsValue   int    `json:"assetsValue"`
}

type SubmissionDao struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
