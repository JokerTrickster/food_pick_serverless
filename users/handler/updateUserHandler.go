package handler

import (
	"main/model/entity"
	_interface "main/model/interface"
	"main/model/request"

	"net/http"

	_env "github.com/JokerTrickster/common/env"
	_error "github.com/JokerTrickster/common/error"
	_jwt "github.com/JokerTrickster/common/jwt"
	_validator "github.com/JokerTrickster/common/validator"

	"github.com/labstack/echo/v4"
)

type UpdateUserHandler struct {
	UseCase _interface.IUpdateUserUseCase
}

func NewUpdateUserHandler(c *echo.Echo, useCase _interface.IUpdateUserUseCase) _interface.IUpdateUserHandler {
	handler := &UpdateUserHandler{
		UseCase: useCase,
	}
	c.PUT("/users/profile", _jwt.TokenChecker(handler.Update))
	return handler
}

func (d *UpdateUserHandler) Update(c echo.Context) error {
	ctx, uID, email := _env.CtxGenerate(c)
	req := &request.ReqUpdateUser{}
	if err := _validator.ValidateReq(c, req); err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}

	entity := entity.UpdateUserEntity{
		UserID: uID,
		Email:  email,
	}
	if req.Birth != "" {
		entity.Birth = req.Birth
	}
	if req.Name != "" {
		entity.Name = req.Name
	}
	if req.Sex != "" {
		entity.Sex = req.Sex
	}
	if req.NewPassword != "" && req.PrevPassword != "" {
		entity.NewPassword = req.NewPassword
		entity.PrevPassword = req.PrevPassword
	}
	if req.Push != nil {
		entity.Push = req.Push
	}

	err := d.UseCase.Update(ctx, &entity)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}

	return c.JSON(http.StatusOK, true)
}
