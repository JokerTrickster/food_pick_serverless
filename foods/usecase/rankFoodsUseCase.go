package usecase

import (
	"context"
	"main/model/entity"
	_interface "main/model/interface"
	"main/model/response"
	"time"

	_redis "github.com/JokerTrickster/common/db/redis"
)

type RankFoodUseCase struct {
	Repository     _interface.IRankFoodRepository
	ContextTimeout time.Duration
}

func NewRankFoodUseCase(repo _interface.IRankFoodRepository, timeout time.Duration) _interface.IRankFoodUseCase {
	return &RankFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *RankFoodUseCase) Rank(c context.Context) (response.ResRankFood, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	currentRanks := make([]*entity.RankFoodRedis, 0)
	var err error
	currentRanks, err = d.Repository.RankTop(ctx)
	if err != nil {
		return response.ResRankFood{}, err
	}
	//현재 랭킹 데이터가 비어 있다면
	if len(currentRanks) == 0 {
		//rdb에서 데이터를 가져와 현재 랭킹에 저장을 한다.
		currentRanks, err = d.Repository.FindRankFoodHistories(ctx)
		if err != nil {
			return response.ResRankFood{}, err
		}
		//현재 랭킹 레디스에 저장한다.

		for _, food := range currentRanks {
			err := d.Repository.IncrementFoodRank(ctx, _redis.RankingKey, food.Name, food.Score)
			if err != nil {
				return response.ResRankFood{}, err
			}
		}
	}

	//이전 데이터가 있다면 랭킹 변동을 계산한다.
	res := response.ResRankFood{}
	for i, food := range currentRanks {
		if i == 10 {
			break
		}
		rank := i + 1
		rankFood := response.RankFood{
			Rank: rank,
			Name: food.Name,
		}

		res.Foods = append(res.Foods, rankFood)
	}

	return res, nil
}
