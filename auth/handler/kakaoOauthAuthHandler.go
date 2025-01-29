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

type KakaoOauthAuthHandler struct {
	UseCase _interface.IKakaoOauthAuthUseCase
}

func NewKakaoOauthAuthHandler(c *echo.Echo, useCase _interface.IKakaoOauthAuthUseCase) _interface.IKakaoOauthAuthHandler {
	handler := &KakaoOauthAuthHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/auth/kakao", handler.KakaoOauth)
	return handler
}

func (d *KakaoOauthAuthHandler) KakaoOauth(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqKakaoOauth{}
	if err := _validator.ValidateReq(c, req); err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}
	res, err := d.UseCase.KakaoOauth(ctx, req)
	if err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}
	return c.JSON(http.StatusOK, res)
}
