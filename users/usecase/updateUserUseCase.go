package usecase

import (
	"context"
	"main/model/entity"
	_interface "main/model/interface"
	"time"
)

type UpdateUserUseCase struct {
	Repository     _interface.IUpdateUserRepository
	ContextTimeout time.Duration
}

func NewUpdateUserUseCase(repo _interface.IUpdateUserRepository, timeout time.Duration) _interface.IUpdateUserUseCase {
	return &UpdateUserUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *UpdateUserUseCase) Update(c context.Context, entity *entity.UpdateUserEntity) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	userDTO, err := CreateUpdateUserDTO(entity)
	if err != nil {
		return err
	}

	if entity.PrevPassword != "" && entity.NewPassword != "" {
		err := d.Repository.CheckPassword(ctx, entity.UserID, entity.PrevPassword)
		if err != nil {
			return err
		}
	}

	//유저 정보를 업데이트 한다.
	err = d.Repository.FindOneAndUpdateUser(ctx, userDTO)
	if err != nil {
		return err
	}
	return nil
}
