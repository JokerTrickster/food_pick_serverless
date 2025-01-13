package request

type ReqSaveFood struct {
	Foods []SaveFood `json:"foods"`
}

type SaveFood struct {
	Name         string   `json:"name" validate:"required"`
	Times        []string `json:"times"`
	Scenarios    []string `json:"scenarios"`
	Themes       []string `json:"themes"`
	Types        []string `json:"types"`
	Amount       string   `json:"amount"`
	Kcal         float64  `json:"kcal"`
	Fat          float64  `json:"fat"`
	Carbohydrate float64  `json:"carbohydrate"`
	Protein      float64  `json:"protein"`
}
