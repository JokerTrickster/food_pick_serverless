package response

type ResV02GoogleOauth struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserID       uint   `json:"userID"`
}
