package repository

import (
	"context"
	"fmt"

	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_env "github.com/JokerTrickster/common/env"
	_error "github.com/JokerTrickster/common/error"
	"gorm.io/gorm"
)

func NewGoogleOauthCallbackAuthRepository(gormDB *gorm.DB) _interface.IGoogleOauthCallbackAuthRepository {
	return &GoogleOauthCallbackAuthRepository{GormDB: gormDB}
}

func (g *GoogleOauthCallbackAuthRepository) DeleteToken(ctx context.Context, uID uint) error {
	token := _mysql.Tokens{
		UserID: uID,
	}
	result := g.GormDB.Model(&token).Where("user_id = ?", uID).Delete(&token)
	if result.Error != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), result.Error.Error(), string(_error.ErrFromInternal))
	}
	return nil
}
func (g *GoogleOauthCallbackAuthRepository) SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error {
	token := _mysql.Tokens{
		UserID:           uID,
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		RefreshExpiredAt: refreshTknExpiredAt,
	}
	result := g.GormDB.Model(&token).Create(&token)
	if result.Error != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), result.Error.Error(), string(_error.ErrFromInternal))
	}
	return nil
}

// 이메일로 체크해서 유저가 있으면 유저 정보 반환하고 없으면 유저를 생성한다.
func (g *GoogleOauthCallbackAuthRepository) FindOneAndUpdateUser(ctx context.Context, email string) (*_mysql.Users, error) {
	user := _mysql.Users{
		Email: email,
	}
	//state = "logout"인 유저 wait으로 변경하고 roomID = 1로 변경 user 객체에 반환
	result := g.GormDB.WithContext(ctx).Model(&user).Where("email = ?", email).First(&user)
	if result.Error == nil {
		// 유저가 존재하면 반환
		return &user, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		// 유저가 존재하지 않으면 생성
		user.Provider = "google"
		user.Birth = "0000-01-01"
		user.Name = "푸드픽"
		user.Push = _env.PtrTrue()
		result = g.GormDB.WithContext(ctx).Model(&user).Create(&user)
		if result.Error != nil {
			return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), fmt.Sprintf("유저 데이터 생성 실패 %v", result.Error), string(_error.ErrFromMysqlDB))
		}
		return &user, nil
	} else {
		// 그 외의 에러 처리
		return nil, _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), fmt.Sprintf("유저 데이터 조회 실패 %v", result.Error), string(_error.ErrFromInternal))
	}
}
