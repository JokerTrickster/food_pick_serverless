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

type ReissueAuthHandler struct {
	UseCase _interface.IReissueAuthUseCase
}

func NewReissueAuthHandler(c *echo.Echo, useCase _interface.IReissueAuthUseCase) _interface.IReissueAuthHandler {
	handler := &ReissueAuthHandler{
		UseCase: useCase,
	}
	c.PUT("/v0.1/auth/token/reissue", handler.Reissue)
	return handler
}

func (d *ReissueAuthHandler) Reissue(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqReissue{}
	if err := _validator.ValidateReq(c, req); err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}
	res, err := d.UseCase.Reissue(ctx, req)
	if err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}
	return c.JSON(http.StatusOK, res)
}
