package handler

import (
	"context"
	_interface "main/model/interface"

	"net/http"

	_error "github.com/JokerTrickster/common/error"
	"github.com/labstack/echo/v4"
)

type MetaFoodHandler struct {
	UseCase _interface.IMetaFoodUseCase
}

func NewMetaFoodHandler(c *echo.Echo, useCase _interface.IMetaFoodUseCase) _interface.IMetaFoodHandler {
	handler := &MetaFoodHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/foods/meta", handler.Meta)
	return handler
}

func (d *MetaFoodHandler) Meta(c echo.Context) error {
	ctx := context.Background()

	res, err := d.UseCase.Meta(ctx)
	if err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}
	// 캐시 히트 여부
	c.Response().Header().Set("X-Cache-Hit", "true")

	return writeCustomJSON(c, http.StatusOK, res)
}
