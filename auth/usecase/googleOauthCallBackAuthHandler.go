package usecase

import (
	"context"
	"fmt"
	"log"
	_interface "main/model/interface"
	"main/model/response"
	"time"

	_google "github.com/JokerTrickster/common/oauth/google"

	"github.com/JokerTrickster/common/db/mysql"
	_jwt "github.com/JokerTrickster/common/jwt"
)

type GoogleOauthCallbackAuthUseCase struct {
	Repository     _interface.IGoogleOauthCallbackAuthRepository
	ContextTimeout time.Duration
}

func NewGoogleOauthCallbackAuthUseCase(repo _interface.IGoogleOauthCallbackAuthRepository, timeout time.Duration) _interface.IGoogleOauthCallbackAuthUseCase {
	return &GoogleOauthCallbackAuthUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *GoogleOauthCallbackAuthUseCase) GoogleOauthCallback(c context.Context, token string) (response.ResGoogleOauthCallback, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	googleService := _google.GetGoogleService()

	// 토큰 검증
	oauthData, err := googleService.Validate(ctx, token)
	fmt.Println(oauthData)
	if err != nil {
		return response.ResGoogleOauthCallback{}, err
	}
	// 유저 생성
	userDTO := CreateGoogleUserDTO(oauthData)
	fmt.Println(userDTO)
	log.Println(userDTO)
	var newUserDTO *mysql.Users
	//유저 존재 체크
	newUserDTO, err = d.Repository.FindOneUser(ctx, userDTO)
	if err != nil {
		return response.ResGoogleOauthCallback{}, err
	}
	if newUserDTO == nil {
		// 유저 정보 insert
		newUserDTO, err = d.Repository.InsertOneUser(ctx, userDTO)
		if err != nil {
			return response.ResGoogleOauthCallback{}, err
		}
	}

	//토큰 생성
	// token create
	accessToken, _, refreshToken, refreshTknExpiredAt, err := _jwt.GenerateToken(newUserDTO.Email, newUserDTO.ID)
	if err != nil {
		return response.ResGoogleOauthCallback{}, err
	}

	// 기존 토큰 제거
	err = d.Repository.DeleteToken(ctx, newUserDTO.ID)
	if err != nil {
		return response.ResGoogleOauthCallback{}, err
	}
	// token db save
	err = d.Repository.SaveToken(ctx, newUserDTO.ID, accessToken, refreshToken, refreshTknExpiredAt)
	if err != nil {
		return response.ResGoogleOauthCallback{}, err
	}

	//response create
	res := response.ResGoogleOauthCallback{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       newUserDTO.ID,
	}

	return res, nil
}
