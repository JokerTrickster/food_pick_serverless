package repository

import (
	"context"
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"

	"gorm.io/gorm"
)

func NewDailyRecommendFoodRepository(gormDB *gorm.DB) _interface.IDailyRecommendFoodRepository {
	return &DailyRecommendFoodRepository{GormDB: gormDB}
}
func (d *DailyRecommendFoodRepository) FindOneFood(ctx context.Context, foodName string) (*_mysql.Foods, error) {
	food := _mysql.Foods{}
	if err := d.GormDB.WithContext(ctx).Model(&food).Where("name = ?", foodName).First(&food).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), foodName), string(_error.ErrFromMysqlDB))
	}
	return &food, nil

}
func (d *DailyRecommendFoodRepository) FindOneFoodImage(ctx context.Context, foodID int) (string, error) {
	foodImage := _mysql.FoodImages{}
	if err := d.GormDB.WithContext(ctx).Model(&foodImage).Where("id = ?", foodID).First(&foodImage).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "food_default.png", nil
		}
		return "", _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), foodID), string(_error.ErrFromMysqlDB))
	}
	return foodImage.Image, nil
}

func (d *DailyRecommendFoodRepository) FindRandomFoods(ctx context.Context, limit int) ([]*_mysql.Foods, error) {
	foods := []*_mysql.Foods{}
	if err := d.GormDB.WithContext(ctx).Model(&_mysql.Foods{}).Order("RAND()").Limit(limit).Find(&foods).Error; err != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), limit), string(_error.ErrFromMysqlDB))
	}
	return foods, nil
}
