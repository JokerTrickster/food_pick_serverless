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

type AllMessageUserHandler struct {
	UseCase _interface.IAllMessageUserUseCase
}

func NewAllMessageUserHandler(c *echo.Echo, useCase _interface.IAllMessageUserUseCase) _interface.IAllMessageUserHandler {
	handler := &AllMessageUserHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/users/message/all", handler.AllMessage)
	return handler
}

func (d *AllMessageUserHandler) AllMessage(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqAllMessageUser{}
	if err := _validator.ValidateReq(c, req); err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	err := d.UseCase.AllMessage(ctx, req)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}

	return c.JSON(http.StatusOK, true)
}
