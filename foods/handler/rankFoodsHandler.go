package handler

import (
	"net/http"

	_interface "main/model/interface"

	_env "github.com/JokerTrickster/common/env"
	_error "github.com/JokerTrickster/common/error"
	"github.com/labstack/echo/v4"
)

type RankFoodHandler struct {
	UseCase _interface.IRankFoodUseCase
}

func NewRankFoodHandler(c *echo.Echo, useCase _interface.IRankFoodUseCase) _interface.IRankFoodHandler {
	handler := &RankFoodHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/foods/rank", handler.Rank)
	return handler
}

func (d *RankFoodHandler) Rank(c echo.Context) error {
	ctx, _, _ := _env.CtxGenerate(c)
	//business logic
	res, err := d.UseCase.Rank(ctx)
	if err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}

	return c.JSON(http.StatusOK, res)
}
