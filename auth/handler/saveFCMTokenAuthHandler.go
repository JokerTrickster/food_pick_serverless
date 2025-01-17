package handler

import (
	_interface "main/model/interface"
	"main/model/request"
	"net/http"

	_env "github.com/JokerTrickster/common/env"
	_error "github.com/JokerTrickster/common/error"
	_jwt "github.com/JokerTrickster/common/jwt"
	_validator "github.com/JokerTrickster/common/validator"
	"github.com/labstack/echo/v4"
)

type SaveFCMTokenAuthHandler struct {
	UseCase _interface.ISaveFCMTokenAuthUseCase
}

func NewSaveFCMTokenAuthHandler(c *echo.Echo, useCase _interface.ISaveFCMTokenAuthUseCase) _interface.ISaveFCMTokenAuthHandler {
	handler := &SaveFCMTokenAuthHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/auth/fcm/token", _jwt.TokenChecker(handler.SaveFCMToken))
	return handler
}

func (d *SaveFCMTokenAuthHandler) SaveFCMToken(c echo.Context) error {
	ctx, uID, _ := _env.CtxGenerate(c)
	req := &request.ReqSaveFCMToken{}
	if err := _validator.ValidateReq(c, req); err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	err := d.UseCase.SaveFCMToken(ctx, uID, req)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	return c.JSON(http.StatusOK, true)
}
