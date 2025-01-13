package response

type ResMetaData struct {
	MetaKeys   []string `json:"metaKeys"`
	MetaKRKeys []string `json:"metaKRKeys"`
	MetaData   MetaData `json:"metaData"`
}

//상황 -> 시간 -> 종륲별 -> 맛 -> 기분/테마별

type MetaData struct {
	Scenarios []Category `json:"scenarios"`
	Times     []Category `json:"times"`
	Types     []Category `json:"types"`
	Themes    []Category `json:"themes"`
}

type Category struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
