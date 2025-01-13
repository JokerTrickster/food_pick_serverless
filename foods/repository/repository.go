package repository

import "gorm.io/gorm"

type DailyRecommendFoodRepository struct {
	GormDB *gorm.DB
}
