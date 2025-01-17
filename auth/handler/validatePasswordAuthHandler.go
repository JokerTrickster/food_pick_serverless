package handler

import (
	"context"
	"main/model/entity"
	_interface "main/model/interface"
	"main/model/request"
	"net/http"

	_error "github.com/JokerTrickster/common/error"
	_validator "github.com/JokerTrickster/common/validator"
	"github.com/labstack/echo/v4"
)

type ValidatePasswordAuthHandler struct {
	UseCase _interface.IValidatePasswordAuthUseCase
}

func NewValidatePasswordAuthHandler(c *echo.Echo, useCase _interface.IValidatePasswordAuthUseCase) _interface.IValidatePasswordAuthHandler {
	handler := &ValidatePasswordAuthHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/auth/password/validate", handler.ValidatePassword)
	return handler
}

func (d *ValidatePasswordAuthHandler) ValidatePassword(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqValidatePassword{}
	if err := _validator.ValidateReq(c, req); err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}

	entity := entity.ValidatePasswordAuthEntity{
		Email:    req.Email,
		Password: req.Password,
		Code:     req.Code,
	}
	err := d.UseCase.ValidatePassword(ctx, entity)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	return c.JSON(http.StatusOK, true)
}
