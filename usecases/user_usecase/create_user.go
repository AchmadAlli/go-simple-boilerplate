package user_usecase

import (
	"context"
)

type CreateUserUsecase interface {
	CreateUser(context.Context, CreateUserInput) (CreateUserOutput, error)
}

type CreateUserPresenter interface {
	PresentCreatedUser() CreateUserOutput
	PresentCreatedUserV2() CreateUserOutputV2
}

type CreateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

type CreateUserOutputV2 struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

type createUserOrchestrator struct {
	userRepo  interface{}
	presenter CreateUserPresenter
}

func NewCrateUserUsecase(
	presenter CreateUserPresenter,
) CreateUserUsecase {
	return &createUserOrchestrator{
		userRepo:  nil,
		presenter: presenter,
	}
}

func (uc createUserOrchestrator) CreateUser(context.Context, CreateUserInput) (CreateUserOutput, error) {
	return CreateUserOutput{}, nil
}
