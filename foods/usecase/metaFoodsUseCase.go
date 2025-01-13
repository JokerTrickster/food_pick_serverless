package usecase

import (
	"context"
	"encoding/json"

	_interface "main/model/interface"
	"main/model/response"
	"time"

	_redis "github.com/JokerTrickster/common/db/redis"
)

type MetaFoodUseCase struct {
	Repository     _interface.IMetaFoodRepository
	ContextTimeout time.Duration
}

func NewMetaFoodUseCase(repo _interface.IMetaFoodRepository, timeout time.Duration) _interface.IMetaFoodUseCase {
	return &MetaFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *MetaFoodUseCase) Meta(c context.Context) (response.ResMetaData, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	// 레디스 조회해서 키값이 있는지 확인
	metaData, err := d.Repository.RedisFindAllMeta(ctx, _redis.FoodMeta)
	if err != nil {
		return response.ResMetaData{}, err
	}
	// 없으면 디비에서 조회해서 레디스에 저장
	if metaData == "" {
		typeDTO, err := d.Repository.FindAllTypeMeta(ctx)
		if err != nil {
			return response.ResMetaData{}, err
		}
		timeDTO, err := d.Repository.FindAllTimeMeta(ctx)
		if err != nil {
			return response.ResMetaData{}, err
		}
		scenarioDTO, err := d.Repository.FindAllScenarioMeta(ctx)
		if err != nil {
			return response.ResMetaData{}, err
		}
		themesDTO, err := d.Repository.FindAllThemesMeta(ctx)
		if err != nil {
			return response.ResMetaData{}, err
		}

		res := CreateResMetaData(typeDTO, timeDTO, scenarioDTO, themesDTO)
		// 3. 조회된 데이터를 Redis에 캐시 (예: 1시간 TTL)
		data, err := json.Marshal(res)
		if err != nil {
			return response.ResMetaData{}, err
		}
		err = d.Repository.RedisSetAllMeta(ctx, _redis.FoodMeta, data)
		if err != nil {
			return response.ResMetaData{}, err
		}

		return res, nil
	}
	var res response.ResMetaData
	if err := json.Unmarshal([]byte(metaData), &res); err != nil {
		return response.ResMetaData{}, err
	}
	return res, nil
}
