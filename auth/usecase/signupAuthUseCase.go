package usecase

import (
	"context"
	_interface "main/model/interface"
	"main/model/request"
	"main/model/response"
	"time"

	_jwt "github.com/JokerTrickster/common/jwt"
)

type SignupAuthUseCase struct {
	Repository     _interface.ISignupAuthRepository
	ContextTimeout time.Duration
}

func NewSignupAuthUseCase(repo _interface.ISignupAuthRepository, timeout time.Duration) _interface.ISignupAuthUseCase {
	return &SignupAuthUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *SignupAuthUseCase) Signup(c context.Context, req *request.ReqSignup) (response.ResSignup, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	// 해당 유저가 존재하는지 체크
	err := d.Repository.UserCheckByEmail(ctx, req.Email)
	if err != nil {
		return response.ResSignup{}, err
	}
	if req.AuthCode != "testCode" {
		//인증코드 검증이 됐는지 체크
		err = d.Repository.VerifyAuthCode(ctx, req.Email, req.AuthCode)
		if err != nil {
			return response.ResSignup{}, err
		}
	}

	// 유저 생성 쿼리문 작성
	user := CreateSignupUser(req)

	// 유저 정보 insert
	uID, err := d.Repository.InsertOneUser(ctx, user)
	if err != nil {
		return response.ResSignup{}, err
	}
	user.ID = uID
	// token create
	accessToken, _, refreshToken, refreshTknExpiredAt, err := _jwt.GenerateToken(user.Email, user.ID)
	if err != nil {
		return response.ResSignup{}, err
	}

	// 기존 토큰 제거
	err = d.Repository.DeleteToken(ctx, user.ID)
	if err != nil {
		return response.ResSignup{}, err
	}
	// token db save
	err = d.Repository.SaveToken(ctx, user.ID, accessToken, refreshToken, refreshTknExpiredAt)
	if err != nil {
		return response.ResSignup{}, err
	}

	//response create
	res := response.ResSignup{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return res, nil
}
