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
	c.GET("/v0.1/auth/email/check", handler.CheckEmail)
	return handler
}

func (d *CheckEmailAuthHandler) CheckEmail(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqCheckEmail{}
	if err := _validator.ValidateReq(c, req); err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}
	err := d.UseCase.CheckEmail(ctx, req.Email)
	if err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}
	return c.JSON(http.StatusOK, true)
}
