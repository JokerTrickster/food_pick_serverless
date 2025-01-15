package usecase

import (
	"context"
	"log"
	_interface "main/model/interface"
	"main/model/request"
	"time"

	"firebase.google.com/go/messaging"
	_error "github.com/JokerTrickster/common/error"
	_firebase "github.com/JokerTrickster/common/firebase"
)

type MessageUserUseCase struct {
	Repository     _interface.IMessageUserRepository
	ContextTimeout time.Duration
}

func NewMessageUserUseCase(repo _interface.IMessageUserRepository, timeout time.Duration) _interface.IMessageUserUseCase {
	return &MessageUserUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *MessageUserUseCase) Message(c context.Context, req *request.ReqMessageUser) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	// 어드민 유저인지 체크한다.
	if req.Role != "foodadmin" {
		return _error.CreateError(ctx, string(_error.ErrBadParameter), _error.Trace(), _error.HandleError("only food admin can send message", req), string(_error.ErrFromClient))
	}
	// 1. 알람 여부를 체크한다.
	alertEnabled, err := d.Repository.FindOneAlarm(ctx, uint(req.UserID))
	if err != nil {
		return err
	}
	if !alertEnabled {
		return _error.CreateError(ctx, string(_error.ErrBadParameter), _error.Trace(), _error.HandleError("user has disabled alert", req), string(_error.ErrFromClient))
	}

	// 2. 푸시 토큰을 가져온다.
	token, err := d.Repository.FindOnePushToken(ctx, uint(req.UserID))
	if err != nil {
		return err
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
		return err
	}
	// 4. 메시지를 저장한다.

	return nil
}
