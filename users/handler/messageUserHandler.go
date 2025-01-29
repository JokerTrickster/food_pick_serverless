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

type MessageUserHandler struct {
	UseCase _interface.IMessageUserUseCase
}

func NewMessageUserHandler(c *echo.Echo, useCase _interface.IMessageUserUseCase) _interface.IMessageUserHandler {
	handler := &MessageUserHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/users/message", handler.Message)
	return handler
}

func (d *MessageUserHandler) Message(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqMessageUser{}
	if err := _validator.ValidateReq(c, req); err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)

	}
	err := d.UseCase.Message(ctx, req)
	if err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)

	}

	return c.JSON(http.StatusOK, true)
}
