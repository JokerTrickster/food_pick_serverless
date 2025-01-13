package repository

import "gorm.io/gorm"

type DailyRecommendFoodRepository struct {
	GormDB *gorm.DB
}

type RecommendFoodRepository struct {
	GormDB *gorm.DB
}