package repository

import (
	"context"
	"main/model/entity"
	_interface "main/model/interface"

	_redis "github.com/JokerTrickster/common/db/redis"
	_error "github.com/JokerTrickster/common/error"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewRankFoodRepository(gormDB *gorm.DB, redisClient *redis.Client) _interface.IRankFoodRepository {
	return &RankFoodRepository{GormDB: gormDB, RedisClient: redisClient}
}

func (g *RankFoodRepository) RankTop(ctx context.Context) ([]*entity.RankFoodRedis, error) {
	//get Ranks foods

	currentRanks, err := g.RedisClient.ZRevRangeWithScores(ctx, _redis.FoodRankingKey, 0, -1).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error()), string(_error.ErrFromRedis))
	}
	result := []*entity.RankFoodRedis{}
	for _, z := range currentRanks {
		rankFood := &entity.RankFoodRedis{
			Name:  z.Member.(string),
			Score: z.Score,
		}
		result = append(result, rankFood)
	}

	return result, nil
}

func (g *RankFoodRepository) FindRankFoodHistories(ctx context.Context) ([]*entity.RankFoodRedis, error) {
	// gorm에서 food_histories 테이블에서 top 10가져오기
	var results []struct {
		name  string
		Count int64
	}
	// SQL 쿼리 실행
	err := g.GormDB.Table("food_histories").
		Select("name, COUNT(name) as count").
		Group("name").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error()), string(_error.ErrFromMysqlDB))
	}

	// 결과에서 음식 이름 추출
	topFoods := make([]*entity.RankFoodRedis, 0)
	for _, r := range results {
		topFoods = append(topFoods, &entity.RankFoodRedis{
			Name:  r.name,
			Score: float64(r.Count),
		})
	}

	return topFoods, nil
}

func (g *RankFoodRepository) IncrementFoodRank(ctx context.Context, redisKey string, foodName string, score float64) error {
	// Increment food Rank in Redis
	_, err := g.RedisClient.ZAdd(ctx, redisKey, redis.Z{Score: score, Member: foodName}).Result()
	if err != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error(), redisKey, foodName, score), string(_error.ErrFromRedis))
	}

	return nil
}
