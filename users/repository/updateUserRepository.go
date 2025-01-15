package repository

import (
	"context"
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"

	"gorm.io/gorm"
)

func NewUpdateUserRepository(gormDB *gorm.DB) _interface.IUpdateUserRepository {
	return &UpdateUserRepository{GormDB: gormDB}
}
func (d *UpdateUserRepository) FindOneAndUpdateUser(ctx context.Context, userDTO *_mysql.Users) error {

	result := d.GormDB.Model(&userDTO).Where("id = ?", userDTO.ID).Updates(&userDTO)
	if result.Error != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(result.Error.Error(), userDTO), string(_error.ErrFromInternal))
	}
	if result.RowsAffected == 0 {
		return _error.CreateError(ctx, string(_error.ErrUserNotFound), _error.Trace(), _error.HandleError(string(_error.ErrUserNotFound), userDTO), string(_error.ErrFromClient))
	}
	return nil
}

func (d *UpdateUserRepository) CheckPassword(ctx context.Context, id uint, prevPassword string) error {
	user := &_mysql.Users{}
	result := d.GormDB.Model(&user).Where("id = ?", id).First(&user)
	if result.Error != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(result.Error.Error(), user), string(_error.ErrFromInternal))
	}
	if user.Password != prevPassword {
		return _error.CreateError(ctx, string(_error.ErrPasswordNotMatch), _error.Trace(), _error.HandleError(string(_error.ErrPasswordNotMatch), user), string(_error.ErrFromClient))
	}
	return nil
}
