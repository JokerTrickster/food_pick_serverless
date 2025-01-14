package _errors

import "errors"

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrInvalidAccessToken     = errors.New("invalid access token")
	ErrUserAlreadyExisted     = errors.New("user already existed")
	ErrCodeOrEmailNotFound    = errors.New("code or email not found")
	ErrInvalidAuthCode        = errors.New("invalid auth code")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
)
