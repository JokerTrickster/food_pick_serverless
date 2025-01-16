package usecase

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"main/model/request"
	"math/rand"
	"net/http"
	"time"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_env "github.com/JokerTrickster/common/env"
	_error "github.com/JokerTrickster/common/error"
	_jwt "github.com/JokerTrickster/common/jwt"
	_oauth "github.com/JokerTrickster/common/oauth"
	_google "github.com/JokerTrickster/common/oauth/google"
)

func CreateGoogleUserDTO(oauthData _oauth.OAuthData) *_mysql.Users {
	return &_mysql.Users{
		Email:    oauthData.Email,
		Password: "",
		Provider: "google",
		Birth:    "1990-01-01",
		Name:     "임시푸드픽",
		Push:     _env.PtrTrue(),
	}
}

// 숫자 5글자 랜덤값 생성
func GeneratePasswordAuthCode() (string, error) {
	rand.Seed(time.Now().UnixNano()) // 난수 시드 초기화
	code := make([]byte, 5)

	for i := 0; i < 5; i++ {
		code[i] = byte(rand.Intn(10) + '0') // 0부터 9까지 랜덤 숫자 생성
	}

	return string(code), nil
}

func CreateTokenDTO(uID uint, accessToken string, accessTknExpiredAt int64, refreshToken string, refreshTknExpiredAt int64) _mysql.Tokens {
	return _mysql.Tokens{
		UserID:           uID,
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		RefreshExpiredAt: refreshTknExpiredAt,
	}
}

func CreateSignupUser(req *request.ReqSignup) _mysql.Users {
	return _mysql.Users{
		Email:    req.Email,
		Password: req.Password,
		Provider: "email",
		Name:     req.Name,
		Birth:    req.Birth,
		Sex:      req.Sex,
		Push:     _env.PtrTrue(),
	}
}

func VerifyAccessAndRefresh(req *request.ReqReissue) error {
	accessTokenUserID, accessTokenEmail, err := _jwt.ParseToken(req.AccessToken)
	if err != nil {
		return err
	}
	refresdhTokenUserID, refreshTokenEmail, err := _jwt.ParseToken(req.RefreshToken)
	if err != nil {
		return err
	}
	if accessTokenUserID != refresdhTokenUserID || accessTokenEmail != refreshTokenEmail {
		return _error.CreateError(context.TODO(), string(_error.ErrBadParameter), _error.Trace(), "access token and refresh token are not matched", string(_error.ErrFromClient))
	}

	if err := _jwt.VerifyToken(req.RefreshToken); err != nil {
		return err
	}
	return nil
}

func GenerateStateOauthCookie(ctx context.Context) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	return state
}

func getGoogleUserInfo(ctx context.Context, accessToken string) ([]byte, error) {
	googleService := _google.GetGoogleService()
	token, err := googleService.ExchangeToken(ctx, accessToken)
	if err != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), fmt.Sprintf("google exchange error %v", err), string(_error.ErrFromInternal))
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), fmt.Sprintf("bad request google access token %v", err), string(_error.ErrFromInternal))
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func CreateSaveFCMTokenDTO(uID uint, req *request.ReqSaveFCMToken) *_mysql.UserTokens {
	return &_mysql.UserTokens{
		UserID: uID,
		Token:  req.Token,
	}
}
