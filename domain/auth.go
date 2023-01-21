package domain

import "errors"

var (
	ErrTokenNotProvided    = errors.New("token not provided")
	ErrUnauthorizedRequest = errors.New("unauthorized request")
	ErrInvalidCredential   = errors.New("invalid credentials")
	ErrForbiddenAccess     = errors.New("forbidden access")
	ErrInvalidToken        = errors.New("invalid token")
	ErrTokenExpired        = errors.New("token expired")
	ErrApiKeyNotProvided   = errors.New("api key not provided")
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"-"`
}

type AuthUsecase interface {
	Authenticate(auth Auth) error
}
