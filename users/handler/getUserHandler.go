package handler

import (
	_interface "main/model/interface"
	"net/http"
	"strconv"

	_env "github.com/JokerTrickster/common/env"
	_error "github.com/JokerTrickster/common/error"
	_jwt "github.com/JokerTrickster/common/jwt"
	"github.com/labstack/echo/v4"
)

type GetUserHandler struct {
	UseCase _interface.IGetUserUseCase
}

func NewGetUserHandler(c *echo.Echo, useCase _interface.IGetUserUseCase) _interface.IGetUserHandler {
	handler := &GetUserHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/users/:userID", _jwt.TokenChecker(handler.Get))
	return handler
}

func (d *GetUserHandler) Get(c echo.Context) error {
	ctx, uID, _ := _env.CtxGenerate(c)
	pathUserID := c.Param("userID")
	puID, _ := strconv.Atoi(pathUserID)
	if pathUserID == "" || uID != uint(puID) {
		return c.JSON(_error.GenerateCustomErrorResponse(http.StatusBadRequest, string(_error.ErrBadParameter), "invalid user id"))
	}

	res, err := d.UseCase.Get(ctx, uID)
	if err != nil {
		return c.JSON(_error.GenerateHTTPErrorResponse(err))
	}

	return c.JSON(http.StatusOK, res)
}
