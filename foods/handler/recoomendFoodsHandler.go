package handler

import (
	"main/model/entity"
	"main/model/request"
	"net/http"

	_interface "main/model/interface"

	_env "github.com/JokerTrickster/common/env"
	_error "github.com/JokerTrickster/common/error"
	_jwt "github.com/JokerTrickster/common/jwt"
	_validator "github.com/JokerTrickster/common/validator"

	"github.com/labstack/echo/v4"
)

type RecommendFoodHandler struct {
	UseCase _interface.IRecommendFoodUseCase
}

func NewRecommendFoodHandler(c *echo.Echo, useCase _interface.IRecommendFoodUseCase) _interface.IRecommendFoodHandler {
	handler := &RecommendFoodHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/foods/recommend", _jwt.TokenChecker(handler.Recommend))
	return handler
}

// DailyRecommendHandler는 음식 추천 핸들러
func (d *RecommendFoodHandler) Recommend(c echo.Context) error {
	ctx, uID, _ := _env.CtxGenerate(c)
	req := &request.ReqRecommendFood{}
	if err := _validator.ValidateReq(c, req); err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}

	//business logic
	entity := entity.RecommendFoodEntity{
		Types:     req.Types,
		Scenarios: req.Scenarios,
		Times:     req.Times,
		Themes:    req.Themes,
		UserID:    uID,
	}
	if req.PreviousAnswer != "" {
		entity.PreviousAnswer = req.PreviousAnswer
	}

	res, err := d.UseCase.Recommend(ctx, entity)
	if err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}
	return writeCustomJSON(c, http.StatusOK, res)
}
