package handler

import (
	"context"
	_interface "main/model/interface"

	"net/http"

	"github.com/labstack/echo/v4"
)

type MetaFoodHandler struct {
	UseCase _interface.IMetaFoodUseCase
}

func NewMetaFoodHandler(c *echo.Echo, useCase _interface.IMetaFoodUseCase) _interface.IMetaFoodHandler {
	handler := &MetaFoodHandler{
		UseCase: useCase,
	}
	c.GET("/foods/meta", handler.Meta)
	return handler
}

func (d *MetaFoodHandler) Meta(c echo.Context) error {
	ctx := context.Background()

	res, err := d.UseCase.Meta(ctx)
	if err != nil {
		return err
	}
	// 캐시 히트 여부
	c.Response().Header().Set("X-Cache-Hit", "true")

	return writeCustomJSON(c, http.StatusOK, res)
}
