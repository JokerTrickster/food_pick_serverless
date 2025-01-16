package repository

import (
	"context"
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"
	"gorm.io/gorm"
)

func NewKakaoOauthAuthRepository(gormDB *gorm.DB) _interface.IKakaoOauthAuthRepository {
	return &KakaoOauthAuthRepository{GormDB: gormDB}
}

func (g *KakaoOauthAuthRepository) DeleteToken(ctx context.Context, uID uint) error {
	token := _mysql.Tokens{
		UserID: uID,
	}
	result := g.GormDB.Model(&token).Where("user_id = ?", uID).Delete(&token)
	if result.Error != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(result.Error.Error(), uID), string(_error.ErrFromInternal))
	}
	return nil
}
func (g *KakaoOauthAuthRepository) SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error {
	token := _mysql.Tokens{
		UserID:           uID,
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		RefreshExpiredAt: refreshTknExpiredAt,
	}
	result := g.GormDB.Model(&token).Create(&token)
	if result.Error != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(result.Error.Error(), uID), string(_error.ErrFromInternal))
	}
	return nil
}

func (g *KakaoOauthAuthRepository) InsertOneUser(ctx context.Context, user *_mysql.Users) (*_mysql.Users, error) {
	result := g.GormDB.WithContext(ctx).Create(&user)
	if result.RowsAffected == 0 {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError("failed user insert", user), string(_error.ErrFromMysqlDB))
	}
	if result.Error != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(result.Error.Error(), user), string(_error.ErrFromMysqlDB))
	}
	return user, nil
}

func (g *KakaoOauthAuthRepository) FindOneUser(ctx context.Context, user *_mysql.Users) (*_mysql.Users, error) {
	var newUser _mysql.Users
	result := g.GormDB.WithContext(ctx).Where("email = ?", user.Email).First(&newUser)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(result.Error.Error(), user), string(_error.ErrFromMysqlDB))
	}
	return &newUser, nil
}
