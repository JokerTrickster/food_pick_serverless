package response

type ResHistoryFood struct {
	Foods []HistoryFood `json:"foods"`
}

type HistoryFood struct {
	Name    string `json:"name"`
	Created string `json:"created"`
}
