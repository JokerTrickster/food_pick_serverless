package usecase

import (
	"context"
	"encoding/json"
	_interface "main/model/interface"
	"main/model/response"
	"time"

	_redis "github.com/JokerTrickster/common/db/redis"
	_jwt "github.com/JokerTrickster/common/jwt"
)

type GuestAuthUseCase struct {
	Repository     _interface.IGuestAuthRepository
	ContextTimeout time.Duration
}

func NewGuestAuthUseCase(repo _interface.IGuestAuthRepository, timeout time.Duration) _interface.IGuestAuthUseCase {
	return &GuestAuthUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *GuestAuthUseCase) Guest(c context.Context) (response.ResGuest, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	var res response.ResGuest

	//레디스 조회한다.
	guestData, err := d.Repository.RedisFindOneGuest(ctx, _redis.FoodGuestKey)
	if guestData == "" {
		// user check
		user, err := d.Repository.FindOneAndUpdateUser(ctx, "test@test.com", "asdasd123")
		if err != nil {
			return response.ResGuest{}, err
		}
		// token create
		accessToken, _, refreshToken, refreshTknExpiredAt, err := _jwt.GenerateToken(user.Email, user.ID)
		if err != nil {
			return response.ResGuest{}, err
		}

		// token db save
		err = d.Repository.SaveToken(ctx, user.ID, accessToken, refreshToken, refreshTknExpiredAt)
		if err != nil {
			return response.ResGuest{}, err
		}

		//response create
		res := response.ResGuest{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
		//레디스에 저장한다.
		// 3. 조회된 데이터를 Redis에 캐시 (예: 1시간 TTL)
		data, err := json.Marshal(res)
		if err != nil {
			return response.ResGuest{}, err
		}
		err = d.Repository.RedisSetOneGuest(ctx, _redis.FoodGuestKey, data)
		if err != nil {
			return response.ResGuest{}, err
		}

		return res, nil
	} else if err != nil {
		return response.ResGuest{}, err
	}
	if err := json.Unmarshal([]byte(guestData), &res); err != nil {
		return response.ResGuest{}, err
	}
	return res, nil
}
