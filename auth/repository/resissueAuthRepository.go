package repository

import (
	"context"
	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"
	_interface "main/model/interface"

	"gorm.io/gorm"
)

func NewReissueAuthRepository(gormDB *gorm.DB) _interface.IReissueAuthRepository {
	return &ReissueAuthRepository{GormDB: gormDB}
}

func (d *ReissueAuthRepository) SaveToken(ctx context.Context, token _mysql.Tokens) error {
	err := d.GormDB.Create(&token).Error
	if err != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), token), string(_error.ErrFromMysqlDB))
	}
	return nil
}

func (d *ReissueAuthRepository) DeleteToken(ctx context.Context, uID uint) error {
	err := d.GormDB.Where("user_id = ?", uID).Delete(&_mysql.Tokens{}).Error
	if err != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), uID), string(_error.ErrFromMysqlDB))
	}
	return nil
}
