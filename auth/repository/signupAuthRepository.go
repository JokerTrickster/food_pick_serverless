package repository

import (
	"context"
	_interface "main/model/interface"
	"time"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"

	"gorm.io/gorm"
)

func NewSignupAuthRepository(gormDB *gorm.DB) _interface.ISignupAuthRepository {
	return &SignupAuthRepository{GormDB: gormDB}
}

func (g *SignupAuthRepository) DeleteToken(ctx context.Context, uID uint) error {
	token := _mysql.Tokens{
		UserID: uID,
	}
	result := g.GormDB.Model(&token).Where("user_id = ?", uID).Delete(&token)
	if result.Error != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(result.Error.Error(), uID), string(_error.ErrFromInternal))
	}
	return nil
}
func (g *SignupAuthRepository) SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error {
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
func (g *SignupAuthRepository) UserCheckByEmail(ctx context.Context, email string) error {
	var user _mysql.Users
	result := g.GormDB.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.RowsAffected == 0 {
		return nil
	} else {
		return _error.CreateError(ctx, string(_error.ErrUserAlreadyExisted), _error.Trace(), _error.HandleError(string(_error.ErrUserAlreadyExisted), email), string(_error.ErrFromClient))
	}
}
func (g *SignupAuthRepository) InsertOneUser(ctx context.Context, user _mysql.Users) (uint, error) {
	result := g.GormDB.WithContext(ctx).Create(&user)
	if result.RowsAffected == 0 {
		return 0, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError("failed user insert", user), string(_error.ErrFromMysqlDB))
	}
	if result.Error != nil {
		return 0, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(result.Error.Error(), user), string(_error.ErrFromMysqlDB))
	}
	return user.ID, nil
}

func (g *SignupAuthRepository) VerifyAuthCode(ctx context.Context, email, code string) error {
	var userAuth _mysql.UserAuths

	tenMinutesAgo := time.Now().Add(-10 * time.Minute).Format("2006-01-02 15:04:05")
	result := g.GormDB.WithContext(ctx).Where("email = ? AND auth_code = ? and created_at >= ? and type = ?", email, code, tenMinutesAgo, "signup").First(&userAuth)
	if result.RowsAffected == 0 {
		return _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(string(_error.ErrInvalidEmailOrPassword), email, code), string(_error.ErrFromClient))
	}
	if result.Error != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(result.Error.Error(), email, code), string(_error.ErrFromMysqlDB))
	}
	return nil
}
