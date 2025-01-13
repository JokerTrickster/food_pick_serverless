package handler

import (
	"context"
	"encoding/json"
	"main/model/response"
	"net/http"
	"time"

	_interface "main/model/interface"

	_redis "github.com/JokerTrickster/common/db/redis"
	"github.com/labstack/echo/v4"
)

type DailyRecommendFoodHandler struct {
	UseCase _interface.IDailyRecommendFoodUseCase
}

func NewDailyRecommendFoodHandler(c *echo.Echo, useCase _interface.IDailyRecommendFoodUseCase) _interface.IDailyRecommendFoodHandler {
	handler := &DailyRecommendFoodHandler{
		UseCase: useCase,
	}
	c.GET("/foods/daily-recommend", handler.DailyRecommend)
	return handler
}

// DailyRecommendHandler는 음식 추천 핸들러
func (d *DailyRecommendFoodHandler) DailyRecommend(c echo.Context) error {
	ctx := context.Background()
	//business logic
	// 레디스 캐시 처리 (aside pattern 사용)
	redisService := _redis.GetRedisService()
	redis, err := redisService.GetClient(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to connect to the redis",
		})
	}
	// 레디스 캐시 조회
	foodData, err := redis.Get(ctx, _redis.DailyKey).Result()
	if foodData == "" {
		// 2. 캐시에 데이터가 없을 경우 UseCase에서 조회
		res, err := d.UseCase.DailyRecommend(ctx)
		if err != nil {
			return err
		}

		// 3. 조회된 데이터를 Redis에 캐시 (예: 1시간 TTL)
		data, err := json.Marshal(res)
		if err != nil {
			return err
		}
		err = redis.Set(ctx, _redis.DailyKey, data, 1*time.Hour).Err()
		if err != nil {
			return err
		}

		// 캐시 히트 여부
		c.Response().Header().Set("X-Cache-Hit", "false")

		// 4. DB에서 조회한 데이터 반환
		return c.JSON(http.StatusOK, res)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch daily recommend foods",
		})
	}

	// 5. 캐시된 데이터가 있을 경우 반환
	var res response.ResDailyRecommendFood
	if err := json.Unmarshal([]byte(foodData), &res); err != nil {
		return err
	}

	// 캐시 히트 여부
	c.Response().Header().Set("X-Cache-Hit", "true")

	// 예제 응답
	return writeCustomJSON(c, http.StatusOK, res)
}

