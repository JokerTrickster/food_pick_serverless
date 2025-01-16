package usecase

import (
	"context"
	_interface "main/model/interface"
	"main/model/request"
	"main/model/response"
	"time"

	_jwt "github.com/JokerTrickster/common/jwt"
)

type SigninAuthUseCase struct {
	Repository     _interface.ISigninAuthRepository
	ContextTimeout time.Duration
}

func NewSigninAuthUseCase(repo _interface.ISigninAuthRepository, timeout time.Duration) _interface.ISigninAuthUseCase {
	return &SigninAuthUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *SigninAuthUseCase) Signin(c context.Context, req *request.ReqSignin) (response.ResSignin, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	// user check
	user, err := d.Repository.FindOneAndUpdateUser(ctx, req.Email, req.Password)
	if err != nil {
		return response.ResSignin{}, err
	}

	// token create
	accessToken, _, refreshToken, refreshTknExpiredAt, err := _jwt.GenerateToken(user.Email, user.ID)
	if err != nil {
		return response.ResSignin{}, err
	}

	// 기존 토큰 제거
	err = d.Repository.DeleteToken(ctx, user.ID)
	if err != nil {
		return response.ResSignin{}, err
	}
	// token db save
	err = d.Repository.SaveToken(ctx, user.ID, accessToken, refreshToken, refreshTknExpiredAt)
	if err != nil {
		return response.ResSignin{}, err
	}

	//response create
	res := response.ResSignin{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user.ID,
	}

	return res, nil
}
