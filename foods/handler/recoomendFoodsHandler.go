package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DailyRecommendHandler는 음식 추천 핸들러
func Recommend(c echo.Context) error {
	log.Println("RecommendHandler invoked")
	// 예제 응답
	return c.JSON(http.StatusOK, map[string]interface{}{
		"recommendations": []string{"Pizza", "Sushi", "Burger"},
	})
}
