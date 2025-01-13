package _errors

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrUserAlreadyExisted = errors.New("user already existed")
	ErrCodeNotFound       = errors.New("code not found")
	ErrGeminiError        = errors.New("gemini error")
	ErrFoodNotFound       = errors.New("food not found")
	ErrServerError        = errors.New("server error")
)
