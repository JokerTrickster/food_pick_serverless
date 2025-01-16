package repository

import (
	"context"
	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"
	_interface "main/model/interface"

	"gorm.io/gorm"
)

func NewRequestSignupAuthRepository(gormDB *gorm.DB) _interface.IRequestSignupAuthRepository {
	return &RequestSignupAuthRepository{GormDB: gormDB}
}

func (g *RequestSignupAuthRepository) FindOneUserByEmail(ctx context.Context, email string) error {
	var user _mysql.Users
	result := g.GormDB.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.RowsAffected != 0 {
		return _error.CreateError(ctx, string(_error.ErrUserAlreadyExisted), _error.Trace(), _error.HandleError(string(_error.ErrUserAlreadyExisted), email), string(_error.ErrFromClient))
	}
	return nil
}

func (d *RequestSignupAuthRepository) InsertAuthCode(ctx context.Context, userAuthDTO _mysql.UserAuths) error {

	//인증 코드를 삽입한다.
	err := d.GormDB.Create(&userAuthDTO).Error
	if err != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), userAuthDTO), string(_error.ErrFromMysqlDB))
	}

	return nil
}

func (d *RequestSignupAuthRepository) DeleteAuthCodeByEmail(ctx context.Context, email string) error {
	err := d.GormDB.Where("email = ? and type = ?", email, "signup").Delete(&_mysql.UserAuths{}).Error
	if err != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), email), string(_error.ErrFromMysqlDB))
	}
	return nil
}
