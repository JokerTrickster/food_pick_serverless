package repository

import (
	"context"
	_errors "main/model/errors"
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"
	"gorm.io/gorm"
)

func NewGetUserRepository(gormDB *gorm.DB) _interface.IGetUserRepository {
	return &GetUserRepository{GormDB: gormDB}
}
func (d *GetUserRepository) FindOneUser(ctx context.Context, uID uint) (*_mysql.Users, error) {

	user := _mysql.Users{}
	result := d.GormDB.Model(&user).Where("id = ?", uID).First(&user)
	if result.Error != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrUserNotFound), _error.Trace(), _error.HandleError(_errors.ErrUserNotFound.Error(), uID), string(_error.ErrFromClient))
	}
	return &user, nil
}
