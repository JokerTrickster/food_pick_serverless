package _interface

import (
	"context"
	"main/model/entity"

	_mysql "github.com/JokerTrickster/common/db/mysql"
)

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
	RedisFindAllMeta(ctx context.Context, key string) (string, error)
	RedisSetAllMeta(ctx context.Context, key string, data []byte) error
}

type IRankFoodRepository interface {
	RankTop(ctx context.Context) ([]*entity.RankFoodRedis, error)
	FindRankFoodHistories(ctx context.Context) ([]*entity.RankFoodRedis, error)
	IncrementFoodRank(ctx context.Context, redisKey string, foodName string, score float64) error
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
	RedisFindAllDailyRecommend(ctx context.Context, key string) (string, error)
	RedisSetAllDailyRecommend(ctx context.Context, key string, data []byte) error
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

type IRecommendFoodRepository interface {
	FindOneRecommendFood(ctx context.Context, query string) (*_mysql.Foods, error)
	FindOneFoodImage(ctx context.Context, id int) (string, error)
	FindOneNutrient(ctx context.Context, foodName string) (*_mysql.Nutrients, error)
	FindOneAndSaveNutrient(ctx context.Context, nutrientDTO *_mysql.Nutrients) (*_mysql.Nutrients, error)
}
