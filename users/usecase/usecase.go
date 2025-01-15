package usecase

import (
	"context"
	"fmt"
	"main/model/entity"
	"main/model/response"
	"time"

	"math/rand"

	_aws "github.com/JokerTrickster/common/aws"
	_mysql "github.com/JokerTrickster/common/db/mysql"

	"gorm.io/gorm"
)

func CreateUpdateUserDTO(entity *entity.UpdateUserEntity) (*_mysql.Users, error) {
	//유저 정보를 업데이트 할 때 사용할 DTO를 생성한다.
	result := &_mysql.Users{
		Model: gorm.Model{
			ID: entity.UserID,
		},
	}
	if entity.Birth != "" {
		result.Birth = entity.Birth
	}
	if entity.Name != "" {
		result.Name = entity.Name
	}
	if entity.Sex != "" {
		result.Sex = entity.Sex
	}
	if entity.Email != "" {
		result.Email = entity.Email
	}
	if entity.PrevPassword != "" && entity.NewPassword != "" {
		result.Password = entity.NewPassword
	}
	if entity.Push != nil {
		result.Push = entity.Push
	}

	return result, nil
}

func CreateResGetUser(user *_mysql.Users) response.ResGetUser {

	//유저 정보를 가져올 때 사용할 DTO를 생성한다.
	res := response.ResGetUser{
		Name:   user.Name,
		Email:  user.Email,
		Sex:    user.Sex,
		Birth:  user.Birth,
		Push:   user.Push,
		UserID: int(user.ID),
	}
	s3Service := _aws.GetS3Service("ap-northeast-2")
	imageUrl, err := s3Service.GetSignedURL(context.TODO(), user.Image, _aws.ImgTypeProfile)
	if err == nil {
		fmt.Println(err)
	}
	res.Image = imageUrl

	return res
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func FileNameGenerateRandom() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
