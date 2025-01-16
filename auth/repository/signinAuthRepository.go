package repository

import (
	"context"
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"

	"gorm.io/gorm"
)

func NewSigninAuthRepository(gormDB *gorm.DB) _interface.ISigninAuthRepository {
	return &SigninAuthRepository{GormDB: gormDB}
}

func (g *SigninAuthRepository) DeleteToken(ctx context.Context, uID uint) error {
	token := _mysql.Tokens{
		UserID: uID,
	}
	result := g.GormDB.Model(&token).Where("user_id = ?", uID).Delete(&token)
	if result.Error != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(result.Error.Error(), uID), string(_error.ErrFromInternal))
	}
	return nil
}
func (g *SigninAuthRepository) SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error {
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

func (g *SigninAuthRepository) FindOneAndUpdateUser(ctx context.Context, email, password string) (_mysql.Users, error) {
	user := _mysql.Users{
		Email: email,
	}
	result := g.GormDB.WithContext(ctx).Model(&user).Where("email = ? and password = ?", email, password).Updates(user)
	if result.Error != nil {
		return _mysql.Users{}, _error.CreateError(ctx, string(_error.ErrUserNotFound), _error.Trace(), _error.HandleError(string(_error.ErrUserNotFound), email, password), string(_error.ErrFromClient))
	}
	if result.RowsAffected == 0 {
		return _mysql.Users{}, _error.CreateError(ctx, string(_error.ErrInvalidEmailOrPassword), _error.Trace(), _error.HandleError(string(_error.ErrInvalidEmailOrPassword), email, password), string(_error.ErrFromClient))
	}
	// 변경된 사용자 정보를 가져옵니다.
	err := g.GormDB.WithContext(ctx).Where("email = ? and provider = ?", email, "email").First(&user).Error
	if err != nil {
		return _mysql.Users{}, _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(err.Error(), email, password), string(_error.ErrFromInternal))
	}
	return user, nil
}
