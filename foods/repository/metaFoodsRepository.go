package repository

import (
	"context"
	_interface "main/model/interface"
	"time"

	_error "github.com/JokerTrickster/common/error"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

func NewMetaFoodRepository(gormDB *gorm.DB, redisClient *redis.Client) _interface.IMetaFoodRepository {
	return &MetaFoodRepository{GormDB: gormDB, RedisClient: redisClient}
}

func (g *MetaFoodRepository) RedisFindAllMeta(ctx context.Context, key string) (string, error) {
	meta, err := g.RedisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error()), string(_error.ErrFromRedis))
	}
	return meta, nil
}
func (g *MetaFoodRepository) RedisSetAllMeta(ctx context.Context, key string, data []byte) error {
	err := g.RedisClient.Set(ctx, key, data, 1*time.Hour).Err()
	if err != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(err.Error()), string(_error.ErrFromRedis))
	}
	return nil
}

func (g *MetaFoodRepository) FindAllTypeMeta(ctx context.Context) ([]_mysql.Types, error) {
	var typeDTO []_mysql.Types
	if err := g.GormDB.WithContext(ctx).Find(&typeDTO).Error; err != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(string(_error.ErrInternalServer)+err.Error()), string(_error.ErrFromMysqlDB))
	}
	return typeDTO, nil
}
func (g *MetaFoodRepository) FindAllTimeMeta(ctx context.Context) ([]_mysql.Times, error) {
	var timeDTO []_mysql.Times
	if err := g.GormDB.WithContext(ctx).Find(&timeDTO).Error; err != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(string(_error.ErrInternalServer)+err.Error()), string(_error.ErrFromMysqlDB))
	}
	return timeDTO, nil
}

func (g *MetaFoodRepository) FindAllScenarioMeta(ctx context.Context) ([]_mysql.Scenarios, error) {
	var scenarioDTO []_mysql.Scenarios
	if err := g.GormDB.WithContext(ctx).Find(&scenarioDTO).Error; err != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(string(_error.ErrInternalServer)+err.Error()), string(_error.ErrFromMysqlDB))
	}
	return scenarioDTO, nil
}

func (g *MetaFoodRepository) FindAllThemesMeta(ctx context.Context) ([]_mysql.Themes, error) {
	var themesDTO []_mysql.Themes
	if err := g.GormDB.WithContext(ctx).Find(&themesDTO).Error; err != nil {
		return nil, _error.CreateError(ctx, string(_error.ErrInternalDB), _error.Trace(), _error.HandleError(string(_error.ErrInternalServer)+err.Error()), string(_error.ErrFromMysqlDB))
	}
	return themesDTO, nil
}
