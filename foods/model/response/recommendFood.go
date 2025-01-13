package response

type ResRecommendFood struct {
	FoodNames []RecommendFood `json:"foodNames"`
}

type RecommendFood struct {
	Name         string  `json:"name"`
	Image        string  `json:"image"`
	Amount       string  `json:"amount"`
	Kcal         float64 `json:"kcal"`
	Carbohydrate float64 `json:"carbohydrate"`
	Protein      float64 `json:"protein"`
	Fat          float64 `json:"fat"`
}
