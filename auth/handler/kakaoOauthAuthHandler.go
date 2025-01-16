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
	c.POST("/auth/kakao", handler.KakaoOauth)
	return handler
}

func (d *KakaoOauthAuthHandler) KakaoOauth(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqKakaoOauth{}
	if err := _validator.ValidateReq(c, req); err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	res, err := d.UseCase.KakaoOauth(ctx, req)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	return c.JSON(http.StatusOK, res)
}
