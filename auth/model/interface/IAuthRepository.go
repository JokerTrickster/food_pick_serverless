package _interface

import (
	"context"
	"main/utils/db/mysql"
)

type ISignupAuthRepository interface {
	UserCheckByEmail(ctx context.Context, email string) error
	InsertOneUser(ctx context.Context, user mysql.Users) (uint, error)
	SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error
	DeleteToken(ctx context.Context, uID uint) error
	VerifyAuthCode(ctx context.Context, email, code string) error
}

type ISigninAuthRepository interface {
	FindOneAndUpdateUser(ctx context.Context, email, password string) (mysql.Users, error)
	SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error
	DeleteToken(ctx context.Context, uID uint) error
}

type ILogoutAuthRepository interface {
	DeleteToken(ctx context.Context, uID uint) error
}

type IReissueAuthRepository interface {
	SaveToken(ctx context.Context, token mysql.Tokens) error
	DeleteToken(ctx context.Context, uID uint) error
}

type IRequestPasswordAuthRepository interface {
	FindOneUserByEmail(ctx context.Context, email string) error
	InsertAuthCode(ctx context.Context, userAuthDTO mysql.UserAuths) error
}

type IValidatePasswordAuthRepository interface {
	CheckAuthCode(ctx context.Context, email, code string) error
	UpdatePassword(ctx context.Context, user mysql.Users) error
	DeleteAuthCode(ctx context.Context, email string) error
}

type ICheckEmailAuthRepository interface {
	CheckEmail(ctx context.Context, email string) error
}
type IGuestAuthRepository interface {
	FindOneAndUpdateUser(ctx context.Context, email, password string) (mysql.Users, error)
	SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error
}

type IGoogleOauthAuthRepository interface {
}

type IGoogleOauthCallbackAuthRepository interface {
	SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error
	DeleteToken(ctx context.Context, uID uint) error
	InsertOneUser(ctx context.Context, user mysql.Users) error
	FindOneAndUpdateUser(ctx context.Context, email string) (*mysql.Users, error)
}

type IV02GoogleOauthCallbackAuthRepository interface {
	SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error
	DeleteToken(ctx context.Context, uID uint) error
	FindOneAndUpdateUser(ctx context.Context, email string) (*mysql.Users, error)
}

type IV02GoogleOauthAuthRepository interface {
	SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error
	DeleteToken(ctx context.Context, uID uint) error
	InsertOneUser(ctx context.Context, user *mysql.Users) (*mysql.Users, error)
	FindOneUser(ctx context.Context, userDTO *mysql.Users) (*mysql.Users, error)
}

type IKakaoOauthAuthRepository interface {
	SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error
	DeleteToken(ctx context.Context, uID uint) error
	InsertOneUser(ctx context.Context, user *mysql.Users) (*mysql.Users, error)
	FindOneUser(ctx context.Context, userDTO *mysql.Users) (*mysql.Users, error)
}

type INaverOauthAuthRepository interface {
	SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error
	DeleteToken(ctx context.Context, uID uint) error
	InsertOneUser(ctx context.Context, user *mysql.Users) (*mysql.Users, error)
	FindOneUser(ctx context.Context, userDTO *mysql.Users) (*mysql.Users, error)
}

type IRequestSignupAuthRepository interface {
	FindOneUserByEmail(ctx context.Context, email string) error
	InsertAuthCode(ctx context.Context, userAuthDTO mysql.UserAuths) error
	DeleteAuthCodeByEmail(ctx context.Context, email string) error
}

type ISaveFCMTokenAuthRepository interface {
	SaveFCMToken(ctx context.Context,userTokenDTO *mysql.UserTokens) error
}
