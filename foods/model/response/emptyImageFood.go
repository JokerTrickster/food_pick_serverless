package response

type ResEmptyImageFood struct {
	Foods []EmptyFood `json:"foods"`
}

type EmptyFood struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
