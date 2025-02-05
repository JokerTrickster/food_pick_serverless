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
	c.POST("/v0.1/auth/google", handler.GoogleOauthCallback)
	return handler
}

func (d *GoogleOauthCallbackAuthHandler) GoogleOauthCallback(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqGoogleOauthCallback{}
	if err := _validator.ValidateReq(c, req); err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	fmt.Println("토큰 ID ,", req.Token)
	res, err := d.UseCase.GoogleOauthCallback(ctx, req.Token)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	return c.JSON(http.StatusOK, res)
}
