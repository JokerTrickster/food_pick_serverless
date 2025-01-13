package repository

import (
	"context"
	"errors"
	_errors "main/model/errors"
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"

	"gorm.io/gorm"
)

func NewRecommendFoodRepository(gormDB *gorm.DB) _interface.IRecommendFoodRepository {
	return &RecommendFoodRepository{GormDB: gormDB}
}

func (d *RecommendFoodRepository) FindOneRecommendFood(ctx context.Context, query string) (*_mysql.Foods, error) {
	food := _mysql.Foods{} // 포인터가 아닌 구조체로 초기화
	result := d.GormDB.WithContext(ctx).Raw(query).Scan(&food)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, _error.CreateError(ctx, string(_error.ErrFoodNotFound), _error.Trace(), "no matching record found", string(_error.ErrFromClient))
		}
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, _error.CreateError(ctx, string(_error.ErrFoodNotFound), _error.Trace(), "no matching record found", string(_error.ErrFromClient))
	}
	return &food, nil // 반환할 때 포인터로 반환
}

func (d *RecommendFoodRepository) FindOneFoodImage(ctx context.Context, id int) (string, error) {
	foodImage := _mysql.FoodImages{}
	err := d.GormDB.WithContext(ctx).Where("id = ?", id).First(&foodImage).Error
	if err != nil {
		return "", _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(_errors.ErrServerError.Error()+err.Error(), id), string(_error.ErrFromMysqlDB))
	}
	return foodImage.Image, nil
}

func (d *RecommendFoodRepository) FindOneNutrient(ctx context.Context, foodName string) (*_mysql.Nutrients, error) {
	nutrient := _mysql.Nutrients{}
	err := d.GormDB.WithContext(ctx).Where("food_name = ?", foodName).First(&nutrient).Error
	if err != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(_errors.ErrServerError.Error()+err.Error(), foodName), string(_error.ErrFromMysqlDB))
	}
	return &nutrient, nil
}

func (d *RecommendFoodRepository) FindOneAndSaveNutrient(ctx context.Context, nutrientDTO *_mysql.Nutrients) (*_mysql.Nutrients, error) {
	nutrient := _mysql.Nutrients{}
	err := d.GormDB.WithContext(ctx).Where("food_name = ?", nutrientDTO.FoodName).First(&nutrient).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 데이터가 없을 경우 저장
			if err := d.GormDB.WithContext(ctx).Create(&nutrientDTO).Error; err != nil {
				return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(_errors.ErrServerError.Error()+err.Error(), nutrientDTO), string(_error.ErrFromMysqlDB))
			}
			return nutrientDTO, nil
		}
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(_errors.ErrServerError.Error()+err.Error(), nutrientDTO), string(_error.ErrFromMysqlDB))
	}
	return &nutrient, nil
}
