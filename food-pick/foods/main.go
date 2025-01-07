package main

import (
	"log"
	handler "main/handler"

	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

var echoLambda *echoadapter.EchoLambda

func Init() {
	e := echo.New()
	// Foods 핸들러 등록
	e.POST("/foods/recommend", handler.Recommend)
	e.GET("/foods/daily-recommend", handler.DailyRecommend)

	// Lambda 실행
	// Echo Lambda 어댑터 초기화
	echoLambda = echoadapter.New(e)
}

func main() {
	Init()
	log.Println("Lambda initialized!")
	lambda.Start(echoLambda.Proxy)
}
