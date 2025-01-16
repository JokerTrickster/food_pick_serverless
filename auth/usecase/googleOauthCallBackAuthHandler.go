package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/model/entity"
	_interface "main/model/interface"
	"main/model/response"
	"time"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_jwt "github.com/JokerTrickster/common/jwt"
)

type GoogleOauthCallbackAuthUseCase struct {
	Repository     _interface.IGoogleOauthCallbackAuthRepository
	ContextTimeout time.Duration
}

func NewGoogleOauthCallbackAuthUseCase(repo _interface.IGoogleOauthCallbackAuthRepository, timeout time.Duration) _interface.IGoogleOauthCallbackAuthUseCase {
	return &GoogleOauthCallbackAuthUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *GoogleOauthCallbackAuthUseCase) GoogleOauthCallback(c context.Context, code string) (response.ResGoogleOauthCallback, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	fmt.Println("code", code)
	data, err := getGoogleUserInfo(ctx, code)
	if err != nil {
		return response.ResGoogleOauthCallback{}, err
	}
	var googleUser entity.GoogleUser
	// JSON 파싱
	if err := json.Unmarshal(data, &googleUser); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	sqlEntity := &entity.GoogleOauthCallbackSQLQuery{
		Email: googleUser.Email,
	}
	var user *_mysql.Users
	//회원 계정이 있으면 통과 없으면 회원 가입
	user, err = d.Repository.FindOneAndUpdateUser(ctx, sqlEntity.Email)
	if err != nil {
		return response.ResGoogleOauthCallback{}, err
	}

	//토큰 생성
	// token create
	accessToken, _, refreshToken, refreshTknExpiredAt, err := _jwt.GenerateToken(user.Email, user.ID)
	if err != nil {
		return response.ResGoogleOauthCallback{}, err
	}

	// 기존 토큰 제거
	err = d.Repository.DeleteToken(ctx, user.ID)
	if err != nil {
		return response.ResGoogleOauthCallback{}, err
	}
	// token db save
	err = d.Repository.SaveToken(ctx, user.ID, accessToken, refreshToken, refreshTknExpiredAt)
	if err != nil {
		return response.ResGoogleOauthCallback{}, err
	}

	//response create
	res := response.ResGoogleOauthCallback{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user.ID,
	}

	return res, nil
}
