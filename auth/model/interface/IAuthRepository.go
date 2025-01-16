package _interface

import (
	"context"

	_mysql "github.com/JokerTrickster/common/db/mysql"
)

type ISignupAuthRepository interface {
	UserCheckByEmail(ctx context.Context, email string) error
	InsertOneUser(ctx context.Context, user _mysql.Users) (uint, error)
	SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error
	DeleteToken(ctx context.Context, uID uint) error
	VerifyAuthCode(ctx context.Context, email, code string) error
}

type ISigninAuthRepository interface {
	FindOneAndUpdateUser(ctx context.Context, email, password string) (_mysql.Users, error)
	SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error
	DeleteToken(ctx context.Context, uID uint) error
}

type ILogoutAuthRepository interface {
	DeleteToken(ctx context.Context, uID uint) error
}

type IReissueAuthRepository interface {
	SaveToken(ctx context.Context, token _mysql.Tokens) error
	DeleteToken(ctx context.Context, uID uint) error
}

type IRequestPasswordAuthRepository interface {
	FindOneUserByEmail(ctx context.Context, email string) error
	InsertAuthCode(ctx context.Context, userAuthDTO _mysql.UserAuths) error
}

type IValidatePasswordAuthRepository interface {
	CheckAuthCode(ctx context.Context, email, code string) error
	UpdatePassword(ctx context.Context, user _mysql.Users) error
	DeleteAuthCode(ctx context.Context, email string) error
}

type ICheckEmailAuthRepository interface {
	CheckEmail(ctx context.Context, email string) error
}
type IGuestAuthRepository interface {
	FindOneAndUpdateUser(ctx context.Context, email, password string) (_mysql.Users, error)
	SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error
	RedisFindOneGuest(ctx context.Context, key string) (string, error)
	RedisSetOneGuest(ctx context.Context, key string, data []byte) error
}


type IGoogleOauthCallbackAuthRepository interface {
	SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error
	DeleteToken(ctx context.Context, uID uint) error
	FindOneAndUpdateUser(ctx context.Context, email string) (*_mysql.Users, error)
}


type IKakaoOauthAuthRepository interface {
	SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error
	DeleteToken(ctx context.Context, uID uint) error
	InsertOneUser(ctx context.Context, user *_mysql.Users) (*_mysql.Users, error)
	FindOneUser(ctx context.Context, userDTO *_mysql.Users) (*_mysql.Users, error)
}

type INaverOauthAuthRepository interface {
	SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error
	DeleteToken(ctx context.Context, uID uint) error
	InsertOneUser(ctx context.Context, user *_mysql.Users) (*_mysql.Users, error)
	FindOneUser(ctx context.Context, userDTO *_mysql.Users) (*_mysql.Users, error)
}

type IRequestSignupAuthRepository interface {
	FindOneUserByEmail(ctx context.Context, email string) error
	InsertAuthCode(ctx context.Context, userAuthDTO _mysql.UserAuths) error
	DeleteAuthCodeByEmail(ctx context.Context, email string) error
}

type ISaveFCMTokenAuthRepository interface {
	SaveFCMToken(ctx context.Context, userTokenDTO *_mysql.UserTokens) error
}
