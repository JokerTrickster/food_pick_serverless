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

type SigninAuthHandler struct {
	UseCase _interface.ISigninAuthUseCase
}

func NewSigninAuthHandler(c *echo.Echo, useCase _interface.ISigninAuthUseCase) _interface.ISigninAuthHandler {
	handler := &SigninAuthHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/auth/signin", handler.Signin)
	return handler
}

func (d *SigninAuthHandler) Signin(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqSignin{}
	if err := _validator.ValidateReq(c, req); err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	res, err := d.UseCase.Signin(ctx, req)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	return c.JSON(http.StatusOK, res)
}
