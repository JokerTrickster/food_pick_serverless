package repository

import (
	"context"
	_interface "main/model/interface"

	_mysql "github.com/JokerTrickster/common/db/mysql"
	_error "github.com/JokerTrickster/common/error"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

func NewGuestAuthRepository(gormDB *gorm.DB, redisClient *redis.Client) _interface.IGuestAuthRepository {
	return &GuestAuthRepository{GormDB: gormDB, RedisClient: redisClient}
}

func (g *GuestAuthRepository) SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error {
	token := _mysql.Tokens{
		UserID:           uID,
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		RefreshExpiredAt: refreshTknExpiredAt,
	}
	result := g.GormDB.Model(&token).Create(&token)
	if result.Error != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(result.Error.Error(), uID, accessToken, refreshToken), string(_error.ErrFromInternal))
	}
	return nil
}

func (g *GuestAuthRepository) FindOneAndUpdateUser(ctx context.Context, email, password string) (_mysql.Users, error) {
	user := _mysql.Users{
		Email: email,
	}

	err := g.GormDB.WithContext(ctx).Where("email = ? and password = ? and provider = ?", email, password, "email").First(&user).Error
	if err != nil {
		return _mysql.Users{}, _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(err.Error(), email, password), string(_error.ErrFromInternal))
	}
	return user, nil
}

func (g *GuestAuthRepository) RedisFindOneGuest(ctx context.Context, key string) (string, error) {
	// redis에서 email로 guestID 찾기
	guestID, err := g.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return "", _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(err.Error(), key), string(_error.ErrFromRedis))
	}
	return guestID, nil
}

func (g *GuestAuthRepository) RedisSetOneGuest(ctx context.Context, key string, data []byte) error {
	// redis에 key로 data 저장
	err := g.RedisClient.Set(ctx, key, data, 0).Err()
	if err != nil {
		return _error.CreateError(ctx, string(_error.ErrInternalServer), _error.Trace(), _error.HandleError(err.Error(), key), string(_error.ErrFromRedis))
	}
	return nil
}
