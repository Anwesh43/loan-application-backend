package daos

type InitiateResponseDao struct {
	ApplicationId string
}

type BalanceSheetResponseDao struct {
	Year          string
	Month         int
	ProfitsOrLoss int
	AssetsValue   int
}

type SubmissionDao struct {
	Status  bool
	Message string
}
