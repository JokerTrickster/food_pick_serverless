package _interface

import (
	"context"
	"main/model/entity"

	_mysql "github.com/JokerTrickster/common/db/mysql"
)

type IRecommendFoodRepository interface {
	SaveRecommendFood(ctx context.Context, foodDTO *_mysql.Foods) (*_mysql.Foods, error)
	FindOneOrCreateFoodImage(ctx context.Context, foodImageDTO *_mysql.FoodImages) (*_mysql.FoodImages, error)
}

type ISelectFoodRepository interface {
	FindOneFood(ctx context.Context, foodDTO *_mysql.Foods) (uint, error)
	InsertOneFoodHistory(ctx context.Context, foodHistoryDTO *_mysql.FoodHistories) error
	IncrementFoodRanking(ctx context.Context, foodName string, score float64) error
}
type IHistoryFoodRepository interface {
	FindAllFoodHistory(ctx context.Context, userID uint) ([]_mysql.FoodHistories, error)
	FindOneFood(ctx context.Context, foodID uint) (*_mysql.Foods, error)
}

type IMetaFoodRepository interface {
	FindAllTypeMeta(ctx context.Context) ([]_mysql.Types, error)
	FindAllTimeMeta(ctx context.Context) ([]_mysql.Times, error)
	FindAllScenarioMeta(ctx context.Context) ([]_mysql.Scenarios, error)
	FindAllThemesMeta(ctx context.Context) ([]_mysql.Themes, error)
}

type IRankingFoodRepository interface {
	RankingTop(ctx context.Context) ([]*entity.RankFoodRedis, error)
	FindRankingFoodHistories(ctx context.Context) ([]*entity.RankFoodRedis, error)
	IncrementFoodRanking(ctx context.Context, redisKey string, foodName string, score float64) error
}

type IImageUploadFoodRepository interface {
	FindOneAndUpdateFoodImages(ctx context.Context, foodName, fileName string) error
}

type IEmptyImageFoodRepository interface {
	FindAllEmptyImageFoods(ctx context.Context) ([]_mysql.FoodImages, error)
}

type IDailyRecommendFoodRepository interface {
	FindOneFood(ctx context.Context, foodName string) (*_mysql.Foods, error)
	FindOneFoodImage(ctx context.Context, foodID int) (string, error)
	FindRandomFoods(ctx context.Context, limit int) ([]*_mysql.Foods, error)
}

type ISaveFoodRepository interface {
	SaveFood(ctx context.Context, foodDTO *_mysql.Foods) (uint, error)
	FindOneOrCreateFoodImage(ctx context.Context, foodImageDTO *_mysql.FoodImages) (*_mysql.FoodImages, error)
	SaveNutrient(ctx context.Context, nutrientDTO *_mysql.Nutrients) error
	FindCategoryIDs(ctx context.Context, categories []string) ([]uint, error)
	SaveFoodCategory(ctx context.Context, foodID uint, categoryIDs []uint) error
}

type ICheckImageUploadFoodRepository interface {
}

type IV1RecommendFoodRepository interface {
	FindOneV1RecommendFood(ctx context.Context, query string) (*_mysql.Foods, error)
	SaveRecommendFood(ctx context.Context, foodDTO *_mysql.Foods) (*_mysql.Foods, error)
	FindOneOrCreateFoodImage(ctx context.Context, foodImageDTO *_mysql.FoodImages) (*_mysql.FoodImages, error)
	CountV1RecommendFood(ctx context.Context, query string) (int, error)
	FindOneFoodImage(ctx context.Context, foodID int) (string, error)
	FindOneNutrient(ctx context.Context, foodName string) (*_mysql.Nutrients, error)
	FindOneAndSaveNutrient(ctx context.Context, nutrientDTO *_mysql.Nutrients) (*_mysql.Nutrients, error)
}

type IV12RecommendFoodRepository interface {
	FindOneV12RecommendFood(ctx context.Context, query string) (*_mysql.Foods, error)
	FindOneFoodImage(ctx context.Context, foodID int) (string, error)
	FindOneNutrient(ctx context.Context, foodName string) (*_mysql.Nutrients, error)
	FindOneAndSaveNutrient(ctx context.Context, nutrientDTO *_mysql.Nutrients) (*_mysql.Nutrients, error)
}
