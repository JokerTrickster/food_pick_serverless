package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DailyRecommendHandler는 음식 추천 핸들러
func DailyRecommend(c echo.Context) error {
	fmt.Println("Test")
	// 예제 응답
	return c.JSON(http.StatusOK, map[string]interface{}{
		"daily": []string{"Pizza", "Sushi", "Burger"},
	})
}
