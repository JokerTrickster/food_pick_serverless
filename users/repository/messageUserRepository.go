package repository

import (
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func NewMessageUserRepository(gormDB *gorm.DB) _interface.IMessageUserRepository {
	return &MessageUserRepository{GormDB: gormDB}
}

func (d *MessageUserRepository) FindOnePushToken(ctx context.Context, uID uint) (string, error) {
	var userToken *_mysql.UserTokens
	err := d.GormDB.Where("user_id = ?", uID).First(&userToken).Error
	if err != nil {
		return "", _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError("error finding push token", err), string(_error.ErrFromMysqlDB))
	}
	return userToken.Token, nil
}

func (d *MessageUserRepository) FindOneAlarm(ctx context.Context, uID uint) (bool, error) {
	var user *_mysql.Users
	err := d.GormDB.Where("id = ?", uID).First(&user).Error
	if err != nil {
		return false, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError("error finding user", err), string(_error.ErrFromMysqlDB))
	}
	return *user.Push, nil
}
