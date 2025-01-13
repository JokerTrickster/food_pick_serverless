package usecase

import (
	"context"
	"main/model/entity"
	"main/model/response"

	_aws "github.com/JokerTrickster/common/aws"

	_interface "main/model/interface"
	"time"
)

type RecommendFoodUseCase struct {
	Repository     _interface.IRecommendFoodRepository
	ContextTimeout time.Duration
}

func NewRecommendFoodUseCase(repo _interface.IRecommendFoodRepository, timeout time.Duration) _interface.IRecommendFoodUseCase {
	return &RecommendFoodUseCase{Repository: repo, ContextTimeout: timeout}
}
func (d *RecommendFoodUseCase) Recommend(c context.Context, e entity.RecommendFoodEntity) (response.ResRecommendFood, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	query := CreateRecommendQuery(e)
	food, err := d.Repository.FindOneRecommendFood(ctx, query)
	if err != nil {
		return response.ResRecommendFood{}, err
	}
	//food image ID로 이미지 URL을 가져온다.
	image, err := d.Repository.FindOneFoodImage(ctx, food.ImageID)
	if err != nil {
		return response.ResRecommendFood{}, err
	}
	s3Service := _aws.GetS3Service("ap-northeast-2")
	imageUrl, err := s3Service.GetSignedURL(ctx, image, _aws.ImgTypeFood)
	if err != nil {
		return response.ResRecommendFood{}, err
	}

	nutrientDTO, err := d.Repository.FindOneNutrient(ctx, food.Name)
	if err != nil {
		return response.ResRecommendFood{}, err
	}
	res := CreateResRecommend(food, imageUrl, nutrientDTO)
	return res, nil
}
