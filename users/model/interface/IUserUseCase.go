package _interface

import (
	"context"
	"main/model/entity"
	"main/model/request"
	"main/model/response"
)

type IGetUserUseCase interface {
	Get(c context.Context, uID uint) (response.ResGetUser, error)
}

type IUpdateUserUseCase interface {
	Update(c context.Context, entity *entity.UpdateUserEntity) error
}

type IDeleteUserUseCase interface {
	Delete(c context.Context, uID uint) error
}
type IMessageUserUseCase interface {
	Message(c context.Context, req *request.ReqMessageUser) error
}

type IUpdateProfileUserUseCase interface {
	UpdateProfile(c context.Context, e *entity.UpdateProfileUserEntity) (response.ResUpdateProfileUser, error)
}

type IAllMessageUserUseCase interface {
	AllMessage(c context.Context, req *request.ReqAllMessageUser) error
}
