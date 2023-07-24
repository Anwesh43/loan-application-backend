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
	Name            string `json:"name"`
	YearEstablished string `json:"year_established"`
	Summary         string `json:"summary"`
}
