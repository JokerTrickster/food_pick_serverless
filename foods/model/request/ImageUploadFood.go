package request

type ReqUpdateProfile struct {
	ProfileID     string               `json:"profileID,omitempty"`
	Name          string               `json:"name,omitempty" example:"Cho"`
	ProfileImage  string               `json:"profileImage,omitempty"`
	PatientID     string               `json:"patient,omitempty"`
	Phone         string               `json:"phone,omitempty"`
}