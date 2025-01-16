package repository

import (
	"context"
	"errors"
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"
	"gorm.io/gorm"
)

func NewValidatePasswordAuthRepository(gormDB *gorm.DB) _interface.IValidatePasswordAuthRepository {
	return &ValidatePasswordAuthRepository{GormDB: gormDB}
}

func (g *ValidatePasswordAuthRepository) CheckAuthCode(ctx context.Context, email, code string) error {
	var userAuth _mysql.UserAuths
	err := g.GormDB.WithContext(ctx).Model(&userAuth).Where("email = ? AND auth_code = ?", email, code).First(&userAuth).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return _error.CreateError(ctx, string(_error.ErrBadParameter), _error.Trace(), _error.HandleError(email, code), string(_error.ErrFromClient))
		} else {
			return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), email, code), string(_error.ErrFromMysqlDB))
		}
	}

	return nil
}

func (g *ValidatePasswordAuthRepository) UpdatePassword(ctx context.Context, user _mysql.Users) error {
	err := g.GormDB.WithContext(ctx).Model(&user).Where("email = ?", user.Email).Update("password", &user.Password).Error
	if err != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), user), string(_error.ErrFromMysqlDB))
	}

	return nil
}

func (g *ValidatePasswordAuthRepository) DeleteAuthCode(ctx context.Context, email string) error {
	userAuth := _mysql.UserAuths{}
	err := g.GormDB.WithContext(ctx).Model(&userAuth).Where("email = ?", email).Delete(&userAuth).Error
	if err != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), email), string(_error.ErrFromMysqlDB))
	}

	return nil
}
