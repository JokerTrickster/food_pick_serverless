package handler

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

// writeCustomJSON은 HTML 이스케이프 없이 JSON 응답을 작성
func writeCustomJSON(c echo.Context, status int, data interface{}) error {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(status)

	encoder := json.NewEncoder(c.Response())
	encoder.SetEscapeHTML(false) // HTML 이스케이프 비활성화
	return encoder.Encode(data)
}
