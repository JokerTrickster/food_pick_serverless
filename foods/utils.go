package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	_aws "github.com/JokerTrickster/common/aws"
// 	_mysql "github.com/JokerTrickster/common/db/mysql"
// 	_jwt "github.com/JokerTrickster/common/jwt"
// )

// // initializeJWT initializes the JWT configuration
// func InitializeJWT() {
// 	if err := _jwt.InitJWT(); err != nil {
// 		log.Printf("Failed to initialize JWT: %v", err)
// 	}
// }

// // initializeDatabase initializes the MySQL database using AWS SSM
// func InitializeDatabase(ctx context.Context) string {
// 	// Define SSM keys for MySQL configuration
// 	ssmKeys := []string{
// 		"dev_common_mysql_host",
// 		"dev_common_mysql_port",
// 		"dev_food_mysql_db",
// 		"dev_food_mysql_user",
// 		"dev_food_mysql_password",
// 	}

// 	// Initialize AWS SSM service
// 	_ = _aws.GetAWSService("ap-northeast-2")
// 	ssmService := _aws.SSMService{}

// 	// Fetch MySQL connection parameters from SSM
// 	dbParams, err := ssmService.AwsSsmGetParams(ctx, ssmKeys)
// 	if err != nil {
// 		log.Printf("Failed to fetch MySQL SSM parameters: %v", err)
// 	}

// 	// Format the MySQL connection string
// 	connectionString := formatConnectionString(dbParams)

// 	// Initialize MySQL service
// 	mysqlService := _mysql.GetMySQLService()
// 	if err := mysqlService.Initialize(ctx, connectionString); err != nil {
// 		log.Printf("Failed to initialize MySQL: %v", err)
// 	}

// 	return connectionString
// }

// // formatConnectionString formats the MySQL connection string
// func formatConnectionString(params []string) string {
// 	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
// 		params[3], // user
// 		params[4], // password
// 		params[0], // host
// 		params[1], // port
// 		params[2], // db
// 	)
// }
