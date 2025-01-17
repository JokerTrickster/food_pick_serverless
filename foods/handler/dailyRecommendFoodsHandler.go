package handler

import (
	"context"
	"net/http"

	_interface "main/model/interface"

	_error "github.com/JokerTrickster/common/error"

	"github.com/labstack/echo/v4"
)

type DailyRecommendFoodHandler struct {
	UseCase _interface.IDailyRecommendFoodUseCase
}

func NewDailyRecommendFoodHandler(c *echo.Echo, useCase _interface.IDailyRecommendFoodUseCase) _interface.IDailyRecommendFoodHandler {
	handler := &DailyRecommendFoodHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/foods/daily-recommend", handler.DailyRecommend)
	return handler
}

// DailyRecommendHandler는 음식 추천 핸들러
func (d *DailyRecommendFoodHandler) DailyRecommend(c echo.Context) error {
	ctx := context.Background()
	//business logic
	res, err := d.UseCase.DailyRecommend(ctx)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}

	// 예제 응답
	return writeCustomJSON(c, http.StatusOK, res)
}
