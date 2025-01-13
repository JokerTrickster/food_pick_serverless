package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DailyRecommendHandler는 음식 추천 핸들러
func Recommend(c echo.Context) error {
	fmt.Println("여기 들어옴?")
	log.Println("RecommendHandler invoked")
	// 예제 응답
	return c.JSON(http.StatusOK, map[string]interface{}{
		"recommendations": []string{"Pizza", "Sushi", "Burger"},
	})
}
