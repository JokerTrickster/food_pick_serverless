package response

type ResRankingFood struct {
	Foods []RankFood `json:"foods"`
}

type RankFood struct {
	Rank int    `json:"rank"`
	Name string `json:"name"`
}
