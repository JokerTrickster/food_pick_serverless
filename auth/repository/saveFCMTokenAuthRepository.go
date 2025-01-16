package repository

import (
	"context"
	"errors"
	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"
	_interface "main/model/interface"

	"gorm.io/gorm"
)

func NewSaveFCMTokenAuthRepository(gormDB *gorm.DB) _interface.ISaveFCMTokenAuthRepository {
	return &SaveFCMTokenAuthRepository{GormDB: gormDB}
}

func (d *SaveFCMTokenAuthRepository) SaveFCMToken(ctx context.Context, userTokenDTO *_mysql.UserTokens) error {
	// 기존에 userID에 해당하는 토큰이 존재하면 업데이트하고 없다면 새로 생성
	var existingToken _mysql.UserTokens
	err := d.GormDB.Where("user_id = ?", userTokenDTO.UserID).First(&existingToken).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 기존 레코드가 없으면 새로 생성
			err = d.GormDB.Create(&userTokenDTO).Error
			if err != nil {
				return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), userTokenDTO), string(_error.ErrFromMysqlDB))
			}
		} else {
			return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), userTokenDTO), string(_error.ErrFromMysqlDB))
		}
	} else {
		// 기존 레코드가 있으면 업데이트
		existingToken.Token = userTokenDTO.Token
		err = d.GormDB.Save(&existingToken).Error
		if err != nil {
			return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), userTokenDTO), string(_error.ErrFromMysqlDB))
		}
	}
	return nil
}
