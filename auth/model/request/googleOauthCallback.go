package request

type ReqGoogleOauthCallback struct {
	Code string `query:"code"`
}
