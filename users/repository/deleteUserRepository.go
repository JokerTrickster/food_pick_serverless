package repository

import (
	"context"
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"

	"gorm.io/gorm"
)

func NewDeleteUserRepository(gormDB *gorm.DB) _interface.IDeleteUserRepository {
	return &DeleteUserRepository{GormDB: gormDB}
}

func (d *DeleteUserRepository) DeleteUser(ctx context.Context, uID uint) error {
	result := d.GormDB.Where("id = ?", uID).Delete(&_mysql.Users{})
	if result.Error != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(result.Error.Error(), uID), string(_error.ErrFromInternal))
	}
	if result.RowsAffected == 0 {
		return _error.CreateError(ctx, string(_error.ErrUserNotFound), _error.Trace(), _error.HandleError(string(_error.ErrUserNotFound), uID), string(_error.ErrFromClient))
	}
	return nil
}
