package handler

import (
	"main/model/entity"
	_interface "main/model/interface"

	"net/http"

	_env "github.com/JokerTrickster/common/env"
	_error "github.com/JokerTrickster/common/error"
	_jwt "github.com/JokerTrickster/common/jwt"
	"github.com/labstack/echo/v4"
)

type UpdateProfileUserHandler struct {
	UseCase _interface.IUpdateProfileUserUseCase
}

func NewUpdateProfileUserHandler(c *echo.Echo, useCase _interface.IUpdateProfileUserUseCase) _interface.IUpdateProfileUserHandler {
	handler := &UpdateProfileUserHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/users/profiles/image", _jwt.TokenChecker(handler.UpdateProfile))
	return handler
}

func (d *UpdateProfileUserHandler) UpdateProfile(c echo.Context) error {
	ctx, uID, _ := _env.CtxGenerate(c)
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	e := &entity.UpdateProfileUserEntity{
		UserID: uID,
		Image:  file,
	}
	res, err := d.UseCase.UpdateProfile(ctx, e)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}

	return c.JSON(http.StatusOK, res)
}
