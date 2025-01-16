package repository

import (
	"context"
	"errors"
	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"
	_errors "main/model/errors"
	_interface "main/model/interface"

	"gorm.io/gorm"
)

func NewRequestPasswordAuthRepository(gormDB *gorm.DB) _interface.IRequestPasswordAuthRepository {
	return &RequestPasswordAuthRepository{GormDB: gormDB}
}

func (g *RequestPasswordAuthRepository) FindOneUserByEmail(ctx context.Context, email string) error {
	var user _mysql.Users
	result := g.GormDB.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.RowsAffected == 0 {
		return _error.CreateError(ctx, string(_error.ErrUserNotFound), _error.Trace(), _error.HandleError(_errors.ErrUserNotFound.Error(), email), string(_error.ErrFromClient))
	}
	return nil
}

// 이메일로 찾아서 있으면 업데이트하고 없으면 삽입한다.
func (d *RequestPasswordAuthRepository) InsertAuthCode(ctx context.Context, userAuthDTO _mysql.UserAuths) error {

	var existingUserAuth _mysql.UserAuths

	// 이메일로 기존 레코드가 있는지 확인
	err := d.GormDB.Where("email = ?", userAuthDTO.Email).First(&existingUserAuth).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 레코드가 없으면 삽입
			err = d.GormDB.Create(&userAuthDTO).Error
			if err != nil {
				return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), userAuthDTO), string(_error.ErrFromMysqlDB))
			}
		} else {
			// 다른 에러가 발생하면 에러 반환
			return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), userAuthDTO), string(_error.ErrFromMysqlDB))
		}
	} else {
		// 레코드가 있으면 업데이트
		userAuthDTO.ID = existingUserAuth.ID
		err = d.GormDB.WithContext(ctx).Model(&userAuthDTO).Where("email = ?", userAuthDTO.Email).Update("auth_code", &userAuthDTO.AuthCode).Error
		if err != nil {
			return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), userAuthDTO), string(_error.ErrFromMysqlDB))
		}
	}

	return nil
}
