package handler

import (
	"context"
	_interface "main/model/interface"
	"main/model/request"
	"net/http"

	_error "github.com/JokerTrickster/common/error"
	_validator "github.com/JokerTrickster/common/validator"
	"github.com/labstack/echo/v4"
)

type CheckEmailAuthHandler struct {
	UseCase _interface.ICheckEmailAuthUseCase
}

func NewCheckEmailAuthHandler(c *echo.Echo, useCase _interface.ICheckEmailAuthUseCase) _interface.ICheckEmailAuthHandler {
	handler := &CheckEmailAuthHandler{
		UseCase: useCase,
	}
	c.GET("/auth/email/check", handler.CheckEmail)
	return handler
}

func (d *CheckEmailAuthHandler) CheckEmail(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqCheckEmail{}
	if err := _validator.ValidateReq(c, req); err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	err := d.UseCase.CheckEmail(ctx, req.Email)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	return c.JSON(http.StatusOK, true)
}
