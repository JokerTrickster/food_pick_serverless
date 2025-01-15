package repository

import (
	"context"
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"

	"gorm.io/gorm"
)

func NewAllMessageUserRepository(gormDB *gorm.DB) _interface.IAllMessageUserRepository {
	return &AllMessageUserRepository{GormDB: gormDB}
}

func (d *AllMessageUserRepository) FindUsersForNotifications(ctx context.Context) ([]*_mysql.Users, error) {
	users := []*_mysql.Users{}
	err := d.GormDB.WithContext(ctx).Where("push = ?", true).Find(&users).Error
	if err != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError("error finding users for notifications", err), string(_error.ErrFromMysqlDB))
	}
	return users, nil
}

func (d *AllMessageUserRepository) FindOnePushToken(ctx context.Context, uID uint) (string, error) {
	var userToken *_mysql.UserTokens
	err := d.GormDB.Where("user_id = ?", uID).First(&userToken).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError("error finding push token", err), string(_error.ErrFromMysqlDB))
	}
	return userToken.Token, nil
}
