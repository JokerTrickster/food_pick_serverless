package repository

import (
	"context"
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"

	"gorm.io/gorm"
)

func NewCheckEmailAuthRepository(gormDB *gorm.DB) _interface.ICheckEmailAuthRepository {
	return &CheckEmailAuthRepository{GormDB: gormDB}
}

func (g *CheckEmailAuthRepository) CheckEmail(ctx context.Context, email string) error {
	user := _mysql.Users{
		Email: email,
	}
	//이메일 중복 체크
	result := g.GormDB.WithContext(ctx).Model(&user).Where("email = ?", email).First(&user)
	if result.Error != nil && result.Error.Error() != gorm.ErrRecordNotFound.Error() {
		return _error.CreateError(ctx, string(_error.ErrUserNotFound), _error.Trace(), _error.HandleError(string(_error.ErrUserNotFound)+result.Error.Error(), email), string(_error.ErrFromClient))
	}

	if result.RowsAffected == 1 {
		return _error.CreateError(ctx, string(_error.ErrUserAlreadyExisted), _error.Trace(), _error.HandleError(string(_error.ErrUserAlreadyExisted), email), string(_error.ErrFromClient))
	}
	return nil
}
