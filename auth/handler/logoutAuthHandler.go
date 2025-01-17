package handler

import (
	_interface "main/model/interface"
	"net/http"

	_env "github.com/JokerTrickster/common/env"
	_error "github.com/JokerTrickster/common/error"
	_jwt "github.com/JokerTrickster/common/jwt"
	"github.com/labstack/echo/v4"
)

type LogoutAuthHandler struct {
	UseCase _interface.ILogoutAuthUseCase
}

func NewLogoutAuthHandler(c *echo.Echo, useCase _interface.ILogoutAuthUseCase) _interface.ILogoutAuthHandler {
	handler := &LogoutAuthHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/auth/logout", _jwt.TokenChecker(handler.Logout))
	return handler
}

func (d *LogoutAuthHandler) Logout(c echo.Context) error {
	ctx, uID, _ := _env.CtxGenerate(c)

	err := d.UseCase.Logout(ctx, uID)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}
	return c.JSON(http.StatusOK, true)
}
