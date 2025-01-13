package main

import (
	"context"
	"fmt"
	"log"
	handler "main/handler"
	"main/repository"
	"main/usecase"
	"time"

	_aws "github.com/JokerTrickster/common/aws"
	_mysql "github.com/JokerTrickster/common/db/mysql"
	_redis "github.com/JokerTrickster/common/db/redis"
	_jwt "github.com/JokerTrickster/common/jwt"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

var echoLambda *echoadapter.EchoLambda

func main() {
	// Foods 서버 초기화
	InitServer()

	// Foods 핸들러 등록
	InitHandler()
	lambda.Start(echoLambda.Proxy)
}

func InitHandler() {
	e := echo.New()

	// DB 초기화
	mysqlService := _mysql.GetMySQLService()

	db, _ := mysqlService.GetGORMDB()
	// 레디스 초기화
	redisService := _redis.GetRedisService()
	redis, _ := redisService.GetClient(context.Background())

	// 핸들러 등록
	handler.NewDailyRecommendFoodHandler(e, usecase.NewDailyRecommendFoodUseCase(repository.NewDailyRecommendFoodRepository(db), 10*time.Second))
	handler.NewRecommendFoodHandler(e, usecase.NewRecommendFoodUseCase(repository.NewRecommendFoodRepository(db), 10*time.Second))
	handler.NewRankFoodHandler(e, usecase.NewRankFoodUseCase(repository.NewRankFoodRepository(db, redis), 10*time.Second))
	handler.NewSelectFoodHandler(e, usecase.NewSelectFoodUseCase(repository.NewSelectFoodRepository(db, redis), 10*time.Second))
	// Echo Lambda 어댑터 초기화
	echoLambda = echoadapter.New(e)
}
func InitServer() {
	// Lambda 실행 컨텍스트 생성
	ctx := context.Background()
	// Initialize JWT
	InitializeJWT()
	// Initialize Database
	InitializeDatabase(ctx)

	// Initialize Redis
	InitializeRedis()

	
}
func InitializeRedis() {
	ctx := context.Background()
	// Define SSM keys for Redis configuration
	ssmKeys := []string{
		"dev_common_redis_host",
		"dev_common_redis_port",
		"dev_food_redis_db",
		"dev_food_redis_password",
		"dev_food_redis_user",
	}

	// Initialize AWS SSM service
	_ = _aws.GetAWSService("ap-northeast-2")
	ssmService := _aws.SSMService{}

	// Fetch Redis connection parameters from SSM
	redisParams, err := ssmService.AwsSsmGetParams(ctx, ssmKeys)
	if err != nil {
		log.Printf("Failed to fetch Redis SSM parameters: %v", err)
	}

	// Format the Redis connection string
	redisConnectionString := formatRedisConnectionString(redisParams)
	fmt.Println(redisConnectionString)
	// Initialize Redis service
	redisService := _redis.GetRedisService()
	if err := redisService.Initialize(ctx, redisConnectionString); err != nil {
		log.Printf("Failed to initialize Redis: %v", err)
	}
	log.Println("Redis Initialized")
}

// initializeJWT initializes the JWT configuration
func InitializeJWT() {
	if err := _jwt.InitJWT(); err != nil {
		log.Printf("Failed to initialize JWT: %v", err)
	}
}

// initializeDatabase initializes the MySQL database using AWS SSM
func InitializeDatabase(ctx context.Context) string {
	// Define SSM keys for MySQL configuration
	ssmKeys := []string{
		"dev_common_mysql_host",
		"dev_common_mysql_port",
		"dev_food_mysql_db",
		"dev_food_mysql_password",
		"dev_food_mysql_user",
	}

	// Initialize AWS SSM service
	_ = _aws.GetAWSService("ap-northeast-2")
	ssmService := _aws.SSMService{}

	// Fetch MySQL connection parameters from SSM
	dbParams, err := ssmService.AwsSsmGetParams(ctx, ssmKeys)
	if err != nil {
		log.Printf("Failed to fetch MySQL SSM parameters: %v", err)
	}

	// Format the MySQL connection string
	connectionString := formatConnectionString(dbParams)
	fmt.Println(connectionString)
	// Initialize MySQL service
	mysqlService := _mysql.GetMySQLService()
	if err := mysqlService.Initialize(ctx, connectionString); err != nil {
		log.Printf("Failed to initialize MySQL: %v", err)
	}

	return connectionString
}

// formatConnectionString formats the MySQL connection string
func formatConnectionString(params []string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		params[4], // password
		params[3], // user
		params[0], // host
		params[1], // port
		params[2], // db
	)
}

func formatRedisConnectionString(params []string) string {
	return fmt.Sprintf("redis://%s:%s@%s:%s/%s",
		params[4], //user
		params[3], //password
		params[0], //host
		params[1], //port
		params[2], //db
	)
}

