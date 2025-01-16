package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type SignupAuthRepository struct {
	GormDB *gorm.DB
}
type SigninAuthRepository struct {
	GormDB *gorm.DB
}

type LogoutAuthRepository struct {
	GormDB *gorm.DB
}

type ReissueAuthRepository struct {
	GormDB *gorm.DB
}

type GoogleOauthAuthRepository struct {
	GormDB *gorm.DB
}


type RequestPasswordAuthRepository struct {
	GormDB *gorm.DB
}

type ValidatePasswordAuthRepository struct {
	GormDB *gorm.DB
}
type CheckEmailAuthRepository struct {
	GormDB *gorm.DB
}

type GuestAuthRepository struct {
	GormDB      *gorm.DB
	RedisClient *redis.Client
}

type GoogleOauthCallbackAuthRepository struct {
	GormDB *gorm.DB
}

type KakaoOauthAuthRepository struct {
	GormDB *gorm.DB
}

type RequestSignupAuthRepository struct {
	GormDB *gorm.DB
}

type SaveFCMTokenAuthRepository struct {
	GormDB *gorm.DB
}
