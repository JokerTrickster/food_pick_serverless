package handler

import (
	"main/model/entity"
	_interface "main/model/interface"
	"main/model/request"

	"net/http"

	_env "github.com/JokerTrickster/common/env"
	_jwt "github.com/JokerTrickster/common/jwt"
	_validator "github.com/JokerTrickster/common/validator"
	"github.com/labstack/echo/v4"
)

type SelectFoodHandler struct {
	UseCase _interface.ISelectFoodUseCase
}

func NewSelectFoodHandler(c *echo.Echo, useCase _interface.ISelectFoodUseCase) _interface.ISelectFoodHandler {
	handler := &SelectFoodHandler{
		UseCase: useCase,
	}
	c.POST("/foods/select", _jwt.TokenChecker(handler.Select))
	return handler
}

func (d *SelectFoodHandler) Select(c echo.Context) error {
	ctx, uID, _ := _env.CtxGenerate(c)
	req := &request.ReqSelectFood{}
	if err := _validator.ValidateReq(c, req); err != nil {
		return err
	}

	//business logic
	e := entity.SelectFoodEntity{
		Types:     req.Types,
		Times:     req.Times,
		Name:      req.Name,
		Themes:    req.Themes,
		Scenarios: req.Scenarios,
		UserID:    uID,
	}

	res, err := d.UseCase.Select(ctx, e)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
