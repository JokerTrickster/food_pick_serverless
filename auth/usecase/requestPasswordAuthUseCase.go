package usecase

import (
	"context"
	"main/model/entity"
	_interface "main/model/interface"

	_aws "github.com/JokerTrickster/common/aws"
	_mysql "github.com/JokerTrickster/common/db/mysql"

	"time"
)

type RequestPasswordAuthUseCase struct {
	Repository     _interface.IRequestPasswordAuthRepository
	ContextTimeout time.Duration
}

func NewRequestPasswordAuthUseCase(repo _interface.IRequestPasswordAuthRepository, timeout time.Duration) _interface.IRequestPasswordAuthUseCase {
	return &RequestPasswordAuthUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *RequestPasswordAuthUseCase) RequestPassword(c context.Context, e entity.RequestPasswordAuthEntity) (string, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	// 1. 이메일 사용자 조회
	err := d.Repository.FindOneUserByEmail(ctx, e.Email)
	if err != nil {
		return "", err
	}

	// 2. 비밀번호 재설정 토큰 생성
	authCode, err := GeneratePasswordAuthCode()

	if err != nil {
		return "", err
	}
	// 3. 토큰 저장
	userAuthDTO := _mysql.UserAuths{
		Email:    e.Email,
		AuthCode: authCode,
		Type:     "password",
	}
	err = d.Repository.InsertAuthCode(ctx, userAuthDTO)
	if err != nil {
		return "", err
	}
	//TODO 4. 추후 이메일 전송
	_aws.EmailSendAuthCode(e.Email, authCode)

	return authCode, nil
}
