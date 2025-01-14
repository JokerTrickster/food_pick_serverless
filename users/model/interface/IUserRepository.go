package _interface

import (
	"context"
	"main/utils/db/mysql"
)

type IGetUserRepository interface {
	FindOneUser(ctx context.Context, uID uint) (*mysql.Users, error)
}

type IUpdateUserRepository interface {
	FindOneAndUpdateUser(ctx context.Context, userDTO *mysql.Users) error
	CheckPassword(ctx context.Context, id uint, prevPassword string) error
}

type IDeleteUserRepository interface {
	DeleteUser(ctx context.Context, uID uint) error
}
type IMessageUserRepository interface {
	FindOnePushToken(ctx context.Context, uID uint) (string, error)
	FindOneAlarm(ctx context.Context, uID uint) (bool, error)
}

type IUpdateProfileUserRepository interface {
	UpdateProfileImage(ctx context.Context, uID uint, fileName string) error
}

type IAllMessageUserRepository interface {
	FindUsersForNotifications(ctx context.Context) ([]*mysql.Users, error)
	FindOnePushToken(ctx context.Context, uID uint) (string, error)
}
