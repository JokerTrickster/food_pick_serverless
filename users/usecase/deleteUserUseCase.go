package usecase

import (
	"context"
	_interface "main/model/interface"
	"time"

	_error "github.com/JokerTrickster/common/error"
)

type DeleteUserUseCase struct {
	Repository     _interface.IDeleteUserRepository
	ContextTimeout time.Duration
}

func NewDeleteUserUseCase(repo _interface.IDeleteUserRepository, timeout time.Duration) _interface.IDeleteUserUseCase {
	return &DeleteUserUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *DeleteUserUseCase) Delete(c context.Context, uID uint) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	if uID == 1 {
		return _error.CreateError(ctx, string(_error.ErrBadParameter), _error.Trace(), _error.HandleError("탈퇴할 수 없습니다.", uID), string(_error.ErrFromClient))
	}
	err := d.Repository.DeleteUser(ctx, uID)
	if err != nil {
		return err
	}

	return nil
}
