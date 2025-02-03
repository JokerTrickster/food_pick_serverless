package usecase

import (
	"context"
	_interface "main/model/interface"
	"main/model/request"
	"time"
)

type SaveFCMTokenAuthUseCase struct {
	Repository     _interface.ISaveFCMTokenAuthRepository
	ContextTimeout time.Duration
}

func NewSaveFCMTokenAuthUseCase(repo _interface.ISaveFCMTokenAuthRepository, timeout time.Duration) _interface.ISaveFCMTokenAuthUseCase {
	return &SaveFCMTokenAuthUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *SaveFCMTokenAuthUseCase) SaveFCMToken(c context.Context, uID uint, req *request.ReqSaveFCMToken) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	// 1. FCM 토큰 저장
	userTokenDTO := CreateSaveFCMTokenDTO(uID, req)
	err := d.Repository.SaveFCMToken(ctx, userTokenDTO)
	if err != nil {
		return err
	}

	return nil
}
