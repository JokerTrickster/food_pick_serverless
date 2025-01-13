package usecase

import (
	"context"
	"encoding/json"
	"main/model/response"

	_aws "github.com/JokerTrickster/common/aws"
	_redis "github.com/JokerTrickster/common/db/redis"

	_interface "main/model/interface"
	"time"
)

type DailyRecommendFoodUseCase struct {
	Repository     _interface.IDailyRecommendFoodRepository
	ContextTimeout time.Duration
}

func NewDailyRecommendFoodUseCase(repo _interface.IDailyRecommendFoodRepository, timeout time.Duration) _interface.IDailyRecommendFoodUseCase {
	return &DailyRecommendFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *DailyRecommendFoodUseCase) DailyRecommend(c context.Context) (response.ResDailyRecommendFood, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	var res response.ResDailyRecommendFood
	// 레디스 조회해서 키값이 있는지 확인
	foodData, err := d.Repository.RedisFindAllDailyRecommend(ctx, _redis.FoodDailyKey)
	if err != nil {
		return response.ResDailyRecommendFood{}, err
	}

	if foodData == "" {
		// 없으면 디비에서 조회해서 레디스에 저장
		// db 에서 랜덤으로 음식 3개를 추천해준다.
		foods, err := d.Repository.FindRandomFoods(ctx, 3)
		if err != nil {
			return response.ResDailyRecommendFood{}, err
		}
		res := response.ResDailyRecommendFood{}
		for _, food := range foods {
			// 음식 이미지를 가져온다.
			foodImage, err := d.Repository.FindOneFoodImage(ctx, food.ImageID)
			if err != nil {
				return response.ResDailyRecommendFood{}, err
			}

			s3Service := _aws.GetS3Service("ap-northeast-2")
			imageUrl, err := s3Service.GetSignedURL(ctx, foodImage, _aws.ImgTypeFood)
			if err != nil {
				return response.ResDailyRecommendFood{}, err
			}
			res.DilayFoods = append(res.DilayFoods, response.DailyRecommendFood{
				Name:  food.Name,
				Image: imageUrl,
			})
		}
		// 3. 조회된 데이터를 Redis에 캐시 (예: 1시간 TTL)
		data, err := json.Marshal(res)
		if err != nil {
			return response.ResDailyRecommendFood{}, err
		}
		err = d.Repository.RedisSetAllDailyRecommend(ctx, _redis.FoodDailyKey, data)
		if err != nil {
			return response.ResDailyRecommendFood{}, err
		}
		return res, nil
	}
	if err := json.Unmarshal([]byte(foodData), &res); err != nil {
		return response.ResDailyRecommendFood{}, err
	}

	return res, nil
}
