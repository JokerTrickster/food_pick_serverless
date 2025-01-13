package usecase

import (
	"context"
	"encoding/json"
	"strings"

	"main/model/entity"
	_interface "main/model/interface"
	"main/model/response"
	"time"

	_aws "github.com/JokerTrickster/common/aws"
	_error "github.com/JokerTrickster/common/error"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type SelectFoodUseCase struct {
	Repository     _interface.ISelectFoodRepository
	ContextTimeout time.Duration
}

func NewSelectFoodUseCase(repo _interface.ISelectFoodRepository, timeout time.Duration) _interface.ISelectFoodUseCase {
	return &SelectFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *SelectFoodUseCase) Select(c context.Context, e entity.SelectFoodEntity) (response.ResSelectFood, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	//db에서 조회한다.
	foodDTO := CreateSelectFoodDTO(e)
	foodID, err := d.Repository.FindOneFood(ctx, foodDTO)
	if err != nil {
		return response.ResSelectFood{}, err
	}
	foodDTO.ID = foodID

	//디비에 저장한다.
	foodHistoryDTO := CreateFoodHistoryDTO(foodID, e.UserID, e.Name)
	if err := d.Repository.InsertOneFoodHistory(ctx, foodHistoryDTO); err != nil {
		return response.ResSelectFood{}, err
	}

	//오늘 날짜 음식 궁합 봐주기 (나중에는 유저 정보로 변경)
	//음식 추천 로직 구현
	ssmService := _aws.SSMService{}
	geminiAPIKey, err := ssmService.AwsSsmGetParam(ctx, "food_gemini_id")
	if err != nil {
		return response.ResSelectFood{}, _error.CreateError(ctx, string(_error.ErrPartner), _error.Trace(), string(_error.ErrGeminiError)+err.Error(), string(_error.ErrFromGemini))
	}
	client, err := genai.NewClient(ctx, option.WithAPIKey(geminiAPIKey))
	if err != nil {
		return response.ResSelectFood{}, _error.CreateError(ctx, string(_error.ErrPartner), _error.Trace(), string(_error.ErrGeminiError)+err.Error(), string(_error.ErrFromGemini))
	}
	model := client.GenerativeModel("gemini-1.5-flash")
	//데이터 가공
	question := CreateSelectFoodQuestion(e)
	resp, err := model.GenerateContent(
		ctx,
		genai.Text("너는 맛있는 요리 음식을 알려주는 전문가이다."),
		genai.Text("오늘 날짜와 음식 이름을 받으면 해당 날짜와 음식 궁합을 알려줘"),
		genai.Text("사람들에게 재미요소로 알려줄려고 한다."),
		genai.Text("최대 글자는 300글자 이내로 답변해주고 건강적으로 사주를 봐줘"),
		genai.Text("예를 들면 2024년 9월 3일 김치찌개와 궁합 \n 날짜와 궁합에 대해서 설명..."),
		genai.Text("응답을 해줄 때 특수문자을 넣어서 응답해주면 안된다."),
		genai.Text("지금부터 질문할게"),
		genai.Text(question),
	)

	if err != nil {
		return response.ResSelectFood{}, _error.CreateError(ctx, string(_error.ErrPartner), _error.Trace(), string(_error.ErrGeminiError)+err.Error(), string(_error.ErrFromGemini))
	}
	// 출력 부분 수정
	res := response.ResSelectFood{}
	if len(resp.Candidates) > 0 {
		marshalResponse, _ := json.MarshalIndent(resp, "", "  ")
		var generateResponse entity.ContentResponse
		if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
			return response.ResSelectFood{}, _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), string(_error.ErrInternalServer)+err.Error(), string(_error.ErrFromInternal))
		}
		for _, cad := range *generateResponse.Candidates {
			if cad.Content != nil {
				cleanedString := strings.Trim(cad.Content.Parts[0], "[] \n")
				res.FoodCompatibility = cleanedString
			}
		}

	} else {
		return response.ResSelectFood{}, _error.CreateError(ctx, string(_error.ErrGeminiError), _error.Trace(), string(_error.ErrFoodNotFound), string(_error.ErrFromGemini))
	}
	//레디스 저장한다.

	if err := d.Repository.IncrementFoodRanking(ctx, foodDTO.Name, 1); err != nil {
		return response.ResSelectFood{}, err
	}

	return res, nil
}
