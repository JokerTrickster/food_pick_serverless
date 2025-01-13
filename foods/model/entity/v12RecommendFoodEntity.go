package entity

type V12RecommendFoodEntity struct {
	Types          string `json:"types"`
	Scenarios      string `json:"scenarios"`
	Times          string `json:"times"`
	Themes         string `json:"themes"`
	UserID         uint   `json:"userID"`
	PreviousAnswer string `json:"previousAnswer"`
}

type V12Content struct {
	Parts []string `json:"Parts"`
	Role  string   `json:"Role"`
}
type V12Candidates struct {
	Content *Content `json:"Content"`
}
type V12ContentResponse struct {
	Candidates *[]Candidates `json:"Candidates"`
}
