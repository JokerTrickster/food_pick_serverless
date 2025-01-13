package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type DailyRecommendFoodRepository struct {
	GormDB *gorm.DB
}

type RecommendFoodRepository struct {
	GormDB *gorm.DB
}
type RankFoodRepository struct {
	GormDB      *gorm.DB
	RedisClient *redis.Client
}
