package response

type ResNaverOauth struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserID       uint   `json:"userID"`
}
