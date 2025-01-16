package repository

import (
	"context"
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"
	"gorm.io/gorm"
)

func NewLogoutAuthRepository(gormDB *gorm.DB) _interface.ILogoutAuthRepository {
	return &LogoutAuthRepository{GormDB: gormDB}
}
func (d *LogoutAuthRepository) DeleteToken(ctx context.Context, uID uint) error {
	token := _mysql.Tokens{
		UserID: uID,
	}
	result := d.GormDB.Model(&token).Where("user_id = ?", uID).Delete(&token)
	if result.Error != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(result.Error.Error(), uID), string(_error.ErrFromInternal))
	}
	return nil
}
