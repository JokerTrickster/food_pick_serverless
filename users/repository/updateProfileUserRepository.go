package repository

import (
	"context"
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"

	"gorm.io/gorm"
)

func NewUpdateProfileUserRepository(gormDB *gorm.DB) _interface.IUpdateProfileUserRepository {
	return &UpdateProfileUserRepository{GormDB: gormDB}
}

func (d *UpdateProfileUserRepository) UpdateProfileImage(ctx context.Context, userID uint, filename string) error {
	user := &_mysql.Users{}
	result := d.GormDB.Model(&user).Where("id = ?", userID).Update("image", filename)
	if result.Error != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(result.Error.Error(), user), string(_error.ErrFromInternal))
	}
	if result.RowsAffected == 0 {
		return _error.CreateError(ctx, string(_error.ErrUserNotFound), _error.Trace(), _error.HandleError(string(_error.ErrUserNotFound), user), string(_error.ErrFromClient))
	}
	return nil
}
