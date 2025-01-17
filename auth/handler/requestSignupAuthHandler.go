package handler

import (
	"context"
	"main/model/entity"
	_interface "main/model/interface"
	"main/model/request"
	"net/http"

	_error "github.com/JokerTrickster/common/error"
	_validator "github.com/JokerTrickster/common/validator"
	"github.com/labstack/echo/v4"
)

type RequestSignupAuthHandler struct {
	UseCase _interface.IRequestSignupAuthUseCase
}

func NewRequestSignupAuthHandler(c *echo.Echo, useCase _interface.IRequestSignupAuthUseCase) _interface.IRequestSignupAuthHandler {
	handler := &RequestSignupAuthHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/auth/signup/request", handler.RequestSignup)
	return handler
}

func (d *RequestSignupAuthHandler) RequestSignup(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqRequestSignup{}
	if err := _validator.ValidateReq(c, req); err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}

	entity := entity.RequestSignupAuthEntity{
		Email: req.Email,
	}
	_, err := d.UseCase.RequestSignup(ctx, entity)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	return c.JSON(http.StatusOK, true)
}
