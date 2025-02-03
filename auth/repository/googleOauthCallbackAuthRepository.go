package repository

import (
	"context"
	"fmt"

	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"
	"gorm.io/gorm"
)

func NewGoogleOauthCallbackAuthRepository(gormDB *gorm.DB) _interface.IGoogleOauthCallbackAuthRepository {
	return &GoogleOauthCallbackAuthRepository{GormDB: gormDB}
}
func (g *GoogleOauthCallbackAuthRepository) InsertOneUser(ctx context.Context, user *_mysql.Users) (*_mysql.Users, error) {
	result := g.GormDB.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), fmt.Sprintf("유저 데이터 생성 실패 %v", result.Error), string(_error.ErrFromMysqlDB))
	}
	return user, nil
}
func (g *GoogleOauthCallbackAuthRepository) FindOneUser(ctx context.Context, user *_mysql.Users) (*_mysql.Users, error) {
	var newUser _mysql.Users
	result := g.GormDB.WithContext(ctx).Where("email = ?", user.Email).First(&newUser)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), fmt.Sprintf("유저 데이터 조회 실패 %v", result.Error), string(_error.ErrFromMysqlDB))
	}
	return &newUser, nil
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
