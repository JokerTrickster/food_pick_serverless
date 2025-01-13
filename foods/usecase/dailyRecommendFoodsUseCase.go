package usecase

import (
	"context"
	"fmt"
	"main/model/response"

	_aws "github.com/JokerTrickster/common/aws"

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
		fmt.Println("imageUrl : ", imageUrl)
		res.DilayFoods = append(res.DilayFoods, response.DailyRecommendFood{
			Name:  food.Name,
			Image: imageUrl,
		})
	}

	return res, nil
}
