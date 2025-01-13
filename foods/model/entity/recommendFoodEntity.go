package entity

type RecommendFoodEntity struct {
	Types          string `json:"types"`
	Scenarios      string `json:"scenarios"`
	Times          string `json:"times"`
	Themes         string `json:"themes"`
	UserID         uint   `json:"userID"`
	PreviousAnswer string `json:"previousAnswer"`
}

type Content struct {
	Parts []string `json:"Parts"`
	Role  string   `json:"Role"`
}
type Candidates struct {
	Content *Content `json:"Content"`
}
type ContentResponse struct {
	Candidates *[]Candidates `json:"Candidates"`
}
