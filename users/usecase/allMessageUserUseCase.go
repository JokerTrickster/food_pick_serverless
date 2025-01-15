package usecase

import (
	"context"
	"log"
	_interface "main/model/interface"
	"main/model/request"
	"time"

	_error "github.com/JokerTrickster/common/error"
	_firebase "github.com/JokerTrickster/common/firebase"

	"firebase.google.com/go/messaging"
)

type AllMessageUserUseCase struct {
	Repository     _interface.IAllMessageUserRepository
	ContextTimeout time.Duration
}

func NewAllMessageUserUseCase(repo _interface.IAllMessageUserRepository, timeout time.Duration) _interface.IAllMessageUserUseCase {
	return &AllMessageUserUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *AllMessageUserUseCase) AllMessage(c context.Context, req *request.ReqAllMessageUser) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	// 어드민 유저인지 체크한다.
	if req.Role != "foodadmin" {
		return _error.CreateError(ctx, string(_error.ErrBadParameter), _error.Trace(), _error.HandleError("only food admin can send message", req), string(_error.ErrFromClient))
	}
	// 1. 알람 여부를 체크한다.
	users, err := d.Repository.FindUsersForNotifications(ctx)
	if err != nil {
		return err
	}
	//TODO 추후 고루틴으로 처리할 예정
	for _, user := range users {

		// 2. 푸시 토큰을 가져온다.
		token, err := d.Repository.FindOnePushToken(ctx, uint(user.ID))
		if err != nil {
			return err
		}
		if token == "" {
			continue
		}
		// 3. 푸시를 보낸다.
		// 메시지 생성
		message := &messaging.Message{
			Token: token,
			Notification: &messaging.Notification{
				Title: req.Title,
				Body:  req.Message,
			},
		}

		msg, err := _firebase.GetFirebaseService().GetClient()
		if err != nil {
			log.Printf("error getting firebase client: %v", err)
		}
		// 메시지 전송
		_, err = msg.Send(ctx, message)
		if err != nil {
			log.Printf("error sending message: %v", err)
		}
	}

	// 4. 메시지를 저장한다.

	return nil
}
