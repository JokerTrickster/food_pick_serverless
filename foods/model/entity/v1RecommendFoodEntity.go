package entity

type V1RecommendFoodEntity struct {
	Types          string `json:"types"`
	Scenarios      string `json:"scenarios"`
	Times          string `json:"times"`
	Themes         string `json:"themes"`
	UserID         uint   `json:"userID"`
	PreviousAnswer string `json:"previousAnswer"`
}

type V1Content struct {
	Parts []string `json:"Parts"`
	Role  string   `json:"Role"`
}
type V1Candidates struct {
	Content *Content `json:"Content"`
}
type V1ContentResponse struct {
	Candidates *[]Candidates `json:"Candidates"`
}
