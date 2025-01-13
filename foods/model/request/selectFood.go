package request

type ReqSelectFood struct {
	Types     string `json:"types" example:"한식"`     // 전체, 양식, 한식, 중식 등
	Scenarios string `json:"scenarios" example:"혼밥"` // 전체 , 혼밥, 가족, 친구들
	Times     string `json:"times" example:"점심"`
	Themes    string `json:"themes" example:"스트레스 해소"`
	Name      string `json:"name" example:"된장찌개"`
}
