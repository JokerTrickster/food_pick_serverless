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

type RequestPasswordAuthHandler struct {
	UseCase _interface.IRequestPasswordAuthUseCase
}

func NewRequestPasswordAuthHandler(c *echo.Echo, useCase _interface.IRequestPasswordAuthUseCase) _interface.IRequestPasswordAuthHandler {
	handler := &RequestPasswordAuthHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/auth/password/request", handler.RequestPassword)
	return handler
}

func (d *RequestPasswordAuthHandler) RequestPassword(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqRequestPassword{}
	if err := _validator.ValidateReq(c, req); err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}

	entity := entity.RequestPasswordAuthEntity{
		Email: req.Email,
	}
	_, err := d.UseCase.RequestPassword(ctx, entity)
	if err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}
	return c.JSON(http.StatusOK, true)
}
