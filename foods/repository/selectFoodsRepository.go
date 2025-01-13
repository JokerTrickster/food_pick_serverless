package repository

import (
	"context"
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_redis "github.com/JokerTrickster/common/db/redis"
	_errors "github.com/JokerTrickster/common/error"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewSelectFoodRepository(gormDB *gorm.DB, redisClient *redis.Client) _interface.ISelectFoodRepository {
	return &SelectFoodRepository{GormDB: gormDB, RedisClient: redisClient}
}

func (g *SelectFoodRepository) FindOneFood(ctx context.Context, foodDTO *_mysql.Foods) (uint, error) {
	food := _mysql.Foods{}
	if err := g.GormDB.WithContext(ctx).Model(&food).Where(foodDTO).First(&food).Error; err != nil {
		return 0, _errors.CreateError(ctx, string(_errors.ErrInternalDB), _errors.Trace(), _errors.HandleError(string(_errors.ErrInternalServer)+err.Error(), foodDTO), string(_errors.ErrFromMysqlDB))
	}
	return food.ID, nil
}
func (g *SelectFoodRepository) InsertOneFoodHistory(ctx context.Context, foodHistoryDTO *_mysql.FoodHistories) error {
	if err := g.GormDB.WithContext(ctx).Create(&foodHistoryDTO).Error; err != nil {
		return _errors.CreateError(ctx, string(_errors.ErrInternalDB), _errors.Trace(), _errors.HandleError(string(_errors.ErrInternalServer)+err.Error(), foodHistoryDTO), string(_errors.ErrFromMysqlDB))
	}

	return nil
}

func (g *SelectFoodRepository) IncrementFoodRanking(ctx context.Context, name string, score float64) error {
	redisKey := _redis.FoodRankingKey
	_, err := g.RedisClient.ZIncrBy(ctx, redisKey, score, name).Result()
	if err != nil {
		return _errors.CreateError(ctx, string(_errors.ErrInternalDB), _errors.Trace(), _errors.HandleError(string(_errors.ErrInternalServer)+err.Error(), name), string(_errors.ErrFromRedis))
	}
	return nil
}
