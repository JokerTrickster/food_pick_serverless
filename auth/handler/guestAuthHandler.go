package handler

import (
	"context"
	_interface "main/model/interface"

	"net/http"

	_error "github.com/JokerTrickster/common/error"
	"github.com/labstack/echo/v4"
)

type GuestAuthHandler struct {
	UseCase _interface.IGuestAuthUseCase
}

func NewGuestAuthHandler(c *echo.Echo, useCase _interface.IGuestAuthUseCase) _interface.IGuestAuthHandler {
	handler := &GuestAuthHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/auth/guest", handler.Guest)
	return handler
}

func (d *GuestAuthHandler) Guest(c echo.Context) error {
	ctx := context.Background()
	// 레디스 캐시 처리

	res, err := d.UseCase.Guest(ctx)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	c.Response().Header().Set("X-Cache-Hit", "true")
	return c.JSON(http.StatusOK, res)
}
