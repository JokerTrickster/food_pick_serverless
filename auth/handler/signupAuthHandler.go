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

type SignupAuthHandler struct {
	UseCase _interface.ISignupAuthUseCase
}

func NewSignupAuthHandler(c *echo.Echo, useCase _interface.ISignupAuthUseCase) _interface.ISignupAuthHandler {
	handler := &SignupAuthHandler{
		UseCase: useCase,
	}
	c.POST("/auth/signup", handler.Signup)
	return handler
}

func (d *SignupAuthHandler) Signup(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqSignup{}
	if err := _validator.ValidateReq(c, req); err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	res, err := d.UseCase.Signup(ctx, req)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	return c.JSON(http.StatusCreated, res)
}
