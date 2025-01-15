package usecase

import (
	"context"
	_interface "main/model/interface"
	"main/model/response"
	"time"
)

type GetUserUseCase struct {
	Repository     _interface.IGetUserRepository
	ContextTimeout time.Duration
}

func NewGetUserUseCase(repo _interface.IGetUserRepository, timeout time.Duration) _interface.IGetUserUseCase {
	return &GetUserUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *GetUserUseCase) Get(c context.Context, uID uint) (response.ResGetUser, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	// 유저 정보를 가져온다.
	userDTO, err := d.Repository.FindOneUser(ctx, uID)
	if err != nil {
		return response.ResGetUser{}, err
	}
	res := CreateResGetUser(userDTO)
	return res, nil
}
