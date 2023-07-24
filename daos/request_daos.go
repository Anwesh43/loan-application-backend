package daos

type InitiateApplicationRequest struct {
	Name            string `json:"name"`
	YearEstablished string `json:"year_established"`
}

type FetchBalanceSheetRequest struct {
	Name string `json:"name"`
	Year string `json:"year"`
}

type SubmitApplicationRequest struct {
	Name            string `json:"name"`
	YearEstablished string `json:"year_established"`
	Summary         string `json:"summary"`
}
