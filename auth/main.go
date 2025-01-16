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
	_firebase "github.com/JokerTrickster/common/firebase"
	_jwt "github.com/JokerTrickster/common/jwt"
	_google "github.com/JokerTrickster/common/oauth/google"
	_kakao "github.com/JokerTrickster/common/oauth/kakao"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

var echoLambda *echoadapter.EchoLambda

func main() {
	// 서버 초기화
	InitServer()

	//  핸들러 등록
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
	fmt.Println(redis)

	// 핸들러 등록
	handler.NewCheckEmailAuthHandler(e, usecase.NewCheckEmailAuthUseCase(repository.NewCheckEmailAuthRepository(db), 10*time.Second))
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

	// Initialize Firebase
	InitializeFirebase()

	// oauth kakao 초기화
	InitializeKakao()

	// oauth google 초기화
	InitializeGoogle()

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

func InitializeFirebase() {
	ctx := context.Background()

	// Initialize AWS SSM service
	_ = _aws.GetAWSService("ap-northeast-2")
	ssmService := _aws.SSMService{}

	// Fetch Firebase configuration parameters from SSM
	serviceKey, err := ssmService.AwsSsmGetParam(ctx, "firebase_service_key")
	if err != nil {
		log.Printf("Failed to fetch Firebase SSM parameters: %v", err)
	}

	// Initialize Firebase service
	firebaseService := _firebase.GetFirebaseService()
	err = firebaseService.Initialize(ctx, serviceKey)
	if err != nil {
		log.Printf("Failed to initialize Firebase: %v", err)
	}
	log.Println("Firebase Initialized")
}

func InitializeKakao() {
	ctx := context.Background()

	// Initialize AWS SSM service
	_ = _aws.GetAWSService("ap-northeast-2")
	ssmService := _aws.SSMService{}
	appID, err := ssmService.AwsSsmGetParam(ctx, "dev_kakao_app_id")
	if err != nil {
		log.Printf("Failed to fetch Kakao SSM parameters: %v", err)
	}

	kakaoService := _kakao.GetKakaoService()
	if err := kakaoService.Initialize(appID); err != nil {
		log.Printf("Failed to initialize Kakao: %v", err)
	}
	log.Println("Kakao Initialized")
}

func InitializeGoogle() {
	ctx := context.Background()

	// Initialize AWS SSM service
	_ = _aws.GetAWSService("ap-northeast-2")
	ssmService := _aws.SSMService{}
	googleClientID, err := ssmService.AwsSsmGetParam(ctx, "food_google_client_id")
	if err != nil {
		log.Printf("Failed to fetch Kakao SSM parameters: %v", err)
	}
	googleClientSecret, err := ssmService.AwsSsmGetParam(ctx, "food_google_client_secret")
	if err != nil {
		log.Printf("Failed to fetch Kakao SSM parameters: %v", err)
	}
	redirectUrl, err := ssmService.AwsSsmGetParam(ctx, "food_google_redirect_url")
	if err != nil {
		log.Printf("Failed to fetch Kakao SSM parameters: %v", err)
	}
	googleIosID, err := ssmService.AwsSsmGetParam(ctx, "food_google_ios_id")
	if err != nil {
		log.Printf("Failed to fetch Kakao SSM parameters: %v", err)
	}
	googleAndID, err := ssmService.AwsSsmGetParam(ctx, "food_google_and_id")
	if err != nil {
		log.Printf("Failed to fetch Kakao SSM parameters: %v", err)
	}

	googleService := _google.GetGoogleService()
	googleService.Initialize(googleClientID, googleClientSecret, redirectUrl, googleIosID, []string{googleAndID})
	log.Println("Google Initialized")
	return
}
