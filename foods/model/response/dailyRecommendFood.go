package response

type ResDailyRecommendFood struct {
	DilayFoods []DailyRecommendFood `json:"dilayFoods"`
}

type DailyRecommendFood struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
