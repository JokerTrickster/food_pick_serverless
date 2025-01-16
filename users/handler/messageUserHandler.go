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
	c.POST("/users/message", handler.Message)
	return handler
}

func (d *MessageUserHandler) Message(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqMessageUser{}
	if err := _validator.ValidateReq(c, req); err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))

	}
	err := d.UseCase.Message(ctx, req)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))

	}

	return c.JSON(http.StatusOK, true)
}
