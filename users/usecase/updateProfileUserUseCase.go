package usecase

import (
	"context"
	"main/model/entity"
	_interface "main/model/interface"
	"main/model/response"
	_aws "github.com/JokerTrickster/common/aws"
	"time"
)

type UpdateProfileUserUseCase struct {
	Repository     _interface.IUpdateProfileUserRepository
	ContextTimeout time.Duration
}

func NewUpdateProfileUserUseCase(repo _interface.IUpdateProfileUserRepository, timeout time.Duration) _interface.IUpdateProfileUserUseCase {
	return &UpdateProfileUserUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *UpdateProfileUserUseCase) UpdateProfile(c context.Context, e *entity.UpdateProfileUserEntity) (response.ResUpdateProfileUser, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	//s3 이미지 업로드 한다.
	filename := FileNameGenerateRandom()
	s3Service := _aws.GetS3Service("ap-northeast-2")
	err := s3Service.UploadImage(ctx, e.Image, filename, _aws.ImgTypeProfile)
	if err != nil {
		return response.ResUpdateProfileUser{}, err
	}

	//유저 정보를 업데이트 한다.
	err = d.Repository.UpdateProfileImage(ctx, e.UserID, filename)
	if err != nil {
		return response.ResUpdateProfileUser{}, err
	}

	//s3 이미지 url을 응답한다.
	url, err := s3Service.GetSignedURL(ctx, filename, _aws.ImgTypeProfile)
	if err != nil {
		return response.ResUpdateProfileUser{}, err
	}
	res := response.ResUpdateProfileUser{
		Image: url,
	}

	return res, nil
}
