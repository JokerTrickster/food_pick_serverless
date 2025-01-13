package response

type ResRecommendFood struct {
	FoodNames []RecommendFood `json:"foodNames"`
}

type RecommendFood struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
