package usecase

import (
	"context"
	"fmt"
	_interface "main/model/interface"
	"main/model/request"
	"main/model/response"
	"time"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_jwt "github.com/JokerTrickster/common/jwt"
	_kakao "github.com/JokerTrickster/common/oauth/kakao"
)

type KakaoOauthAuthUseCase struct {
	Repository     _interface.IKakaoOauthAuthRepository
	ContextTimeout time.Duration
}

func NewKakaoOauthAuthUseCase(repo _interface.IKakaoOauthAuthRepository, timeout time.Duration) _interface.IKakaoOauthAuthUseCase {
	return &KakaoOauthAuthUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *KakaoOauthAuthUseCase) KakaoOauth(c context.Context, req *request.ReqKakaoOauth) (response.ResKakaoOauth, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	// 토큰 검증
	kakao := _kakao.GetKakaoService()
	oauthData, err := kakao.Validate(ctx, req.Token)
	if err != nil {
		fmt.Println(err)
		return response.ResKakaoOauth{}, err
	}

	// 유저 생성
	userDTO := CreateGoogleUserDTO(oauthData)
	var newUserDTO *_mysql.Users
	//유저 존재 체크
	newUserDTO, err = d.Repository.FindOneUser(ctx, userDTO)
	if err != nil {
		return response.ResKakaoOauth{}, err
	}
	if newUserDTO == nil {
		// 유저 정보 insert
		newUserDTO, err = d.Repository.InsertOneUser(ctx, userDTO)
		if err != nil {
			return response.ResKakaoOauth{}, err
		}
	}

	// token create
	accessToken, _, refreshToken, refreshTknExpiredAt, err := _jwt.GenerateToken(newUserDTO.Email, newUserDTO.ID)
	if err != nil {
		return response.ResKakaoOauth{}, err
	}

	// 기존 토큰 제거
	err = d.Repository.DeleteToken(ctx, newUserDTO.ID)
	if err != nil {
		return response.ResKakaoOauth{}, err
	}
	// token db save
	err = d.Repository.SaveToken(ctx, newUserDTO.ID, accessToken, refreshToken, refreshTknExpiredAt)
	if err != nil {
		return response.ResKakaoOauth{}, err
	}
	res := response.ResKakaoOauth{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       newUserDTO.ID,
	}

	return res, nil
}
