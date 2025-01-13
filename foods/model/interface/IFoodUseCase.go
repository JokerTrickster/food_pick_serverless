package _interface

import (
	"context"
	"main/model/entity"
	"main/model/request"
	"main/model/response"
)

type ISelectFoodUseCase interface {
	Select(c context.Context, entity entity.SelectFoodEntity) (response.ResSelectFood, error)
}

type IHistoryFoodUseCase interface {
	History(c context.Context, userID uint) (response.ResHistoryFood, error)
}

type IMetaFoodUseCase interface {
	Meta(c context.Context) (response.ResMetaData, error)
}

type IRankingFoodUseCase interface {
	Ranking(c context.Context) (response.ResRankingFood, error)
}

type IImageUploadFoodUseCase interface {
	ImageUpload(c context.Context, e entity.ImageUploadFoodEntity) error
}
type IEmptyImageFoodUseCase interface {
	EmptyImage(c context.Context) (response.ResEmptyImageFood, error)
}

type IDailyRecommendFoodUseCase interface {
	DailyRecommend(c context.Context) (response.ResDailyRecommendFood, error)
}

type ISaveFoodUseCase interface {
	Save(c context.Context, req *request.ReqSaveFood) error
}

type ICheckImageUploadFoodUseCase interface {
	CheckImageUpload(c context.Context, req *request.ReqCheckImageUploadFood) error
}

type IRecommendFoodUseCase interface {
	Recommend(c context.Context, entity entity.RecommendFoodEntity) (response.ResRecommendFood, error)
}
