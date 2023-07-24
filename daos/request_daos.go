package daos

type InitiateApplicationRequest struct {
	Name            string
	YearEstablished string
}

type FetchBalanceSheetRequest struct {
	Name string
	Year string
}

type SubmitApplicationRequest struct {
	Name            string
	YearEstablished string
	Summary         string
}
