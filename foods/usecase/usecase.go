package usecase

import (
	"fmt"
	"main/model/entity"
	"main/model/request"
	"main/model/response"

	_mysql "github.com/JokerTrickster/common/db/mysql"

	"strconv"
	"strings"
	"time"
)

func CreateRecommendFoodImageDTO(entity entity.RecommendFoodEntity, foodName string) *_mysql.FoodImages {

	return &_mysql.FoodImages{
		Name:  foodName,
		Image: "food_default.png",
	}
}

func CreateDailyRecommendFoodQuestion() string {
	today := time.Now().Format("2006-01-02")
	question := fmt.Sprintf("오늘 날짜 %s와 궁합이 좋은 음식 3개 추천해줘 음식 이름만 추천해줘", today)
	return question
}
func CreateResEmptyImageFood(foodImages []_mysql.FoodImages) response.ResEmptyImageFood {
	var res response.ResEmptyImageFood
	for _, f := range foodImages {
		emptyFood := response.EmptyFood{
			ID:   f.ID,
			Name: f.Name,
		}
		res.Foods = append(res.Foods, emptyFood)
	}
	return res
}
func CreateSelectFoodQuestion(e entity.SelectFoodEntity) string {
	today := time.Now().Format("2006-01-02")
	question := fmt.Sprintf("오늘 날짜 %s 와 %s 음식 궁합을 알려줘", today, e.Name)
	return question
}

func CreateSelectFoodDTO(entity entity.SelectFoodEntity) *_mysql.Foods {
	return &_mysql.Foods{
		Name: entity.Name,
	}
}
func CreateFoodHistoryDTO(foodID, userID uint, name string) *_mysql.FoodHistories {
	return &_mysql.FoodHistories{
		FoodID: int(foodID),
		UserID: int(userID),
	}
}

func CreateRecommendFoodDTO(entity entity.RecommendFoodEntity, foodName string, foodImageID int) *_mysql.Foods {
	return &_mysql.Foods{
		Name:    foodName,
		ImageID: foodImageID,
	}

}

func SplitAndRemoveEmpty(s string) []string {
	// 문자열의 연속된 공백을 단일 공백으로 치환하고 앞뒤 공백 제거
	trimmedString := strings.TrimSpace(s)
	// 공백을 기준으로 문자열 분할
	words := strings.Fields(trimmedString)
	return words
}

func CreateRecommendFoodQuestion(entity entity.RecommendFoodEntity) string {
	var reqType string
	if entity.Types == "" || entity.Types == "전체" {
		reqType = "전체 음식"
	} else {
		reqType = entity.Types
	}
	var reqScenario string
	if entity.Scenarios == "" || entity.Scenarios == "전체" {
		reqScenario = "어떤 상황이든"
	} else {
		reqScenario = entity.Scenarios
	}
	var reqTime string
	if entity.Times == "" || entity.Times == "전체" {
		reqTime = "아무때나"
	} else {
		reqTime = entity.Times
	}
	var reqTheme string
	if entity.Themes == "" || entity.Themes == "전체" {
		reqTheme = "아무 테마"
	} else {
		reqTheme = entity.Themes
	}

	questionType := fmt.Sprintf("어떤 종류의 음식 :  %s \n", reqType)
	questionScenario := fmt.Sprintf("누구와 함께 : %s \n", reqScenario)
	questionTime := fmt.Sprintf("언제 : %s \n", reqTime)
	questionTheme := fmt.Sprintf("어떤 테마 : %s \n", reqTheme)
	today := time.Now().Format("2006-01-02")
	question := fmt.Sprintf("%s와 어울리는 %s, %s, %s, %s 음식 이름 1개만 추천해줘 설명 필요없고 이름만 추천해줘", today, questionType, questionScenario, questionTime, questionTheme)
	if entity.PreviousAnswer != "" {
		question += fmt.Sprintf("이전에 추천받은 음식은 제외하고 알려줘 이전 추천 음식 이름 : %s", entity.PreviousAnswer)
	}

	return question
}

func CreateFoodDTOList(req *request.ReqSaveFood) []*_mysql.Foods {
	var foods []*_mysql.Foods
	for _, f := range req.Foods {
		food := _mysql.Foods{
			Name: f.Name,
		}
		foods = append(foods, &food)
	}
	return foods
}

func CreateSaveFoodImageDTO(food request.SaveFood) *_mysql.FoodImages {
	return &_mysql.FoodImages{
		Name:  food.Name,
		Image: "food_default.png",
	}
}

func CreateSaveFoodDTO(food request.SaveFood, foodImageID int) *_mysql.Foods {
	return &_mysql.Foods{
		Name:    food.Name,
		ImageID: foodImageID,
	}
}

func CreateSaveNutrientDTO(food request.SaveFood) *_mysql.Nutrients {
	return &_mysql.Nutrients{
		FoodName:     food.Name,
		Kcal:         food.Kcal,
		Fat:          food.Fat,
		Carbohydrate: food.Carbohydrate,
		Protein:      food.Protein,
		Amount:       food.Amount,
	}
}

func CreateV1RecommendFoodQuestion(entity entity.V1RecommendFoodEntity) string {
	var reqType string
	if entity.Types == "" || entity.Types == "전체" {
		reqType = "전체 음식(한식, 중식, 일식, 양식, 분식, 베트남 음식, 인도 음식, 퓨전 요리)"
	} else {
		reqType = entity.Types
	}
	var reqScenario string
	if entity.Scenarios == "" || entity.Scenarios == "전체" {
		reqScenario = "어떤 상황이든(가족, 연인, 친구, 혼자, 회식)"
	} else {
		reqScenario = entity.Scenarios
	}
	var reqTime string
	if entity.Times == "" || entity.Times == "전체" {
		reqTime = "아무때나(아침, 점심, 저녁, 야식)"
	} else {
		reqTime = entity.Times
	}
	var reqTheme string
	if entity.Themes == "" || entity.Themes == "전체" {
		reqTheme = "아무 테마(스트레스 해소, 해장, 피로 회복, 다이어트, 제철 음식)"
	} else {
		reqTheme = entity.Themes
	}

	questionType := fmt.Sprintf("어떤 종류의 음식 :  %s \n", reqType)
	questionScenario := fmt.Sprintf("누구와 함께 : %s \n", reqScenario)
	questionTime := fmt.Sprintf("언제 : %s \n", reqTime)
	questionTheme := fmt.Sprintf("어떤 테마 : %s \n", reqTheme)
	question := fmt.Sprintf("%s, %s, %s, %s 음식 이름 추천해줘 음식 이름에 공백이 있으면 안된다.", questionType, questionScenario, questionTime, questionTheme)
	if entity.PreviousAnswer != "" {
		question += fmt.Sprintf("이전에 추천받은 음식은 제외하고 알려줘 이전 추천 음식 이름 : %s", entity.PreviousAnswer)
	}
	fmt.Println(question)

	return question
}

func CreateRecommendQuery(entity entity.V1RecommendFoodEntity) string {
	var query string = "SELECT * FROM foods WHERE "
	if entity.Types != "" {
		query += fmt.Sprintf("type_id = (SELECT id FROM types WHERE name = '%s') AND ", entity.Types)
	}
	if entity.Times != "" {
		query += fmt.Sprintf("time_id = (SELECT id FROM times WHERE name = '%s') AND ", entity.Times)
	}
	if entity.Scenarios != "" {
		query += fmt.Sprintf("scenario_id = (SELECT id FROM scenarios WHERE name = '%s') AND ", entity.Scenarios)
	}
	if entity.Themes != "" {
		query += fmt.Sprintf("theme_id = (SELECT id FROM themes WHERE name = '%s') AND ", entity.Themes)
	}

	query = strings.TrimSuffix(query, " AND ")
	return query
}
func CreateV12RecommendQuery(entity entity.V12RecommendFoodEntity) string {
	// Base query
	query := `
		SELECT DISTINCT f.name, f.id, f.image_id
		FROM foods f`

	// 조건별로 JOIN 추가
	if entity.Types != "" {
		query += `
		JOIN food_categories fc_type ON f.id = fc_type.food_id
		JOIN categories c_type ON fc_type.category_id = c_type.id AND c_type.name = '` + entity.Types + `' 
		AND c_type.type_id = (SELECT id FROM category_types WHERE name = 'type')`
	}

	if entity.Scenarios != "" {
		query += `
		JOIN food_categories fc_scenario ON f.id = fc_scenario.food_id
		JOIN categories c_scenario ON fc_scenario.category_id = c_scenario.id AND c_scenario.name = '` + entity.Scenarios + `' 
		AND c_scenario.type_id = (SELECT id FROM category_types WHERE name = 'scenario')`
	}

	if entity.Times != "" {
		query += `
		JOIN food_categories fc_time ON f.id = fc_time.food_id
		JOIN categories c_time ON fc_time.category_id = c_time.id AND c_time.name = '` + entity.Times + `' 
		AND c_time.type_id = (SELECT id FROM category_types WHERE name = 'time')`
	}

	if entity.Themes != "" {
		query += `
		JOIN food_categories fc_theme ON f.id = fc_theme.food_id
		JOIN categories c_theme ON fc_theme.category_id = c_theme.id AND c_theme.name = '` + entity.Themes + `' 
		AND c_theme.type_id = (SELECT id FROM category_types WHERE name = 'theme')`
	}

	// PreviousAnswer 처리
	if entity.PreviousAnswer != "" {
		previous := "'" + strings.Join(strings.Split(entity.PreviousAnswer, ","), "','") + "'"
		query += `
		WHERE f.name NOT IN (` + previous + `)`
	}

	// 랜덤 정렬 및 결과 제한
	query += `
	ORDER BY RAND()
	LIMIT 1`

	fmt.Println(query) // 디버그용 출력
	return query
}

func CreateV1RecommendFoodImageDTO(entity entity.V1RecommendFoodEntity, foodName string) *_mysql.FoodImages {

	return &_mysql.FoodImages{
		Name:  foodName,
		Image: "food_default.png",
	}
}

func CreateV1RecommendFoodDTO(entity entity.V1RecommendFoodEntity, foodName string, foodImageID int) *_mysql.Foods {
	return &_mysql.Foods{
		Name:    foodName,
		ImageID: foodImageID,
	}

}

func CreateRes1Recommend(food *_mysql.Foods, imageUrl string, nutrientDTO *_mysql.Nutrients) response.ResV1RecommendFood {
	res := response.ResV1RecommendFood{}
	foodRes := response.V1RecommendFood{
		Name:         food.Name,
		Image:        imageUrl,
		Amount:       nutrientDTO.Amount,
		Kcal:         nutrientDTO.Kcal,
		Carbohydrate: nutrientDTO.Carbohydrate,
		Protein:      nutrientDTO.Protein,
		Fat:          nutrientDTO.Fat,
	}
	res.FoodNames = append(res.FoodNames, foodRes)
	return res
}

func CreateRes12Recommend(food *_mysql.Foods, imageUrl string, nutrientDTO *_mysql.Nutrients) response.ResV12RecommendFood {
	res := response.ResV12RecommendFood{}
	foodRes := response.V12RecommendFood{
		Name:         food.Name,
		Image:        imageUrl,
		Amount:       nutrientDTO.Amount,
		Kcal:         nutrientDTO.Kcal,
		Carbohydrate: nutrientDTO.Carbohydrate,
		Protein:      nutrientDTO.Protein,
		Fat:          nutrientDTO.Fat,
	}
	res.FoodNames = append(res.FoodNames, foodRes)
	return res
}

// 응답 파싱 함수
func ParseFoodResponse(foodResponse []string) (string, *_mysql.Nutrients, error) {
	// 공백으로 구분하여 분리
	foodName := foodResponse[0]
	amount := foodResponse[1]
	kcal := foodResponse[2]
	carbohydrate := foodResponse[3]
	protein := foodResponse[4]
	fat := foodResponse[5]

	nutrition := &_mysql.Nutrients{
		FoodName: foodName,
		Amount:   amount,
		Kcal: func() float64 {
			kcalFloat, err := strconv.ParseFloat(kcal, 64)
			if err != nil {
				return 0
			}
			return kcalFloat
		}(),
		Carbohydrate: func() float64 {
			carbohydrateFloat, err := strconv.ParseFloat(carbohydrate, 64)
			if err != nil {
				return 0
			}
			return carbohydrateFloat
		}(),
		Protein: func() float64 {
			proteinFloat, err := strconv.ParseFloat(protein, 64)
			if err != nil {
				return 0
			}
			return proteinFloat
		}(),
		Fat: func() float64 {
			fatFloat, err := strconv.ParseFloat(fat, 64)
			if err != nil {
				return 0
			}
			return fatFloat
		}(),
	}

	return foodName, nutrition, nil
}

func CreateCategory(req request.SaveFood) []string {
	var category []string
	category = append(category, req.Types...)
	category = append(category, req.Times...)
	category = append(category, req.Scenarios...)
	category = append(category, req.Themes...)
	return category
}
