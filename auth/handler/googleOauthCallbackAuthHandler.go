package handler

import (
	"context"
	"fmt"
	_interface "main/model/interface"
	"net/http"

	"main/model/request"

	_error "github.com/JokerTrickster/common/error"
	_validator "github.com/JokerTrickster/common/validator"
	"github.com/labstack/echo/v4"
)

type GoogleOauthCallbackAuthHandler struct {
	UseCase _interface.IGoogleOauthCallbackAuthUseCase
}

func NewGoogleOauthCallbackAuthHandler(c *echo.Echo, useCase _interface.IGoogleOauthCallbackAuthUseCase) _interface.IGoogleOauthCallbackAuthHandler {
	handler := &GoogleOauthCallbackAuthHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/auth/google/callback", handler.GoogleOauthCallback)
	return handler
}

func (d *GoogleOauthCallbackAuthHandler) GoogleOauthCallback(c echo.Context) error {
	fmt.Println("GoogleOauthCallback")
	ctx := context.Background()
	req := &request.ReqGoogleOauthCallback{}
	if err := _validator.ValidateReq(c, req); err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}

	res, err := d.UseCase.GoogleOauthCallback(ctx, req.Code)
	fmt.Println(res, err)
	if err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}

	return c.JSON(http.StatusOK, res)
}
