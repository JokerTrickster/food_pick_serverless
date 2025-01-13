package request

type ReqCheckImageUploadFood struct {
	FailedFoodList  []string `json:"failedFoodList"`
	SuccessFoodList []string `json:"successFoodList"`
}
