package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// User 구조체 정의
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Handler 함수
func getUserHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// 요청으로부터 유저 ID 추출
	fmt.Println("테스트 입니다")
	fmt.Println("이거 왜 안찍힘?")
	log.Println("이거는 찍히나?")
	// JSON 응답 생성
	responseBody, err := json.Marshal(User{ID: "1", Name: "홍길동"})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"message": "Error generating response: %v"}`, err),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
	}, nil
}
func main() {
	lambda.Start(getUserHandler)
}
