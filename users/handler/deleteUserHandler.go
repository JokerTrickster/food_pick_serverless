package handler

import (
	_interface "main/model/interface"
	"net/http"

	_env "github.com/JokerTrickster/common/env"
	_error "github.com/JokerTrickster/common/error"
	_jwt "github.com/JokerTrickster/common/jwt"

	"github.com/labstack/echo/v4"
)

type DeleteUserHandler struct {
	UseCase _interface.IDeleteUserUseCase
}

func NewDeleteUserHandler(c *echo.Echo, useCase _interface.IDeleteUserUseCase) _interface.IDeleteUserHandler {
	handler := &DeleteUserHandler{
		UseCase: useCase,
	}
	c.DELETE("/v0.1/users", _jwt.TokenChecker(handler.Delete))
	return handler
}

func (d *DeleteUserHandler) Delete(c echo.Context) error {
	ctx, uID, _ := _env.CtxGenerate(c)

	err := d.UseCase.Delete(ctx, uID)
	if err != nil {
		httpCode, resError := _error.GenerateHTTPErrorResponse(err)
		// 반드시 에러를 반환
		return echo.NewHTTPError(httpCode, resError)
	}

	return c.JSON(http.StatusOK, true)
}
