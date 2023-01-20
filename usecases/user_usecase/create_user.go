package user_usecase

import (
	"context"

	"github.com/achmadAlli/go-simple-boilerplate/domain"
)

type CreateUserUsecase interface {
	CreateUser(context.Context, CreateUserInput) (CreateUserOutput, error)
}

type CreateUserPresenter interface {
	PresentCreatedUser(user domain.User) CreateUserOutput
	PresentCreatedUserV2(user domain.User) CreateUserOutputV2
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
	CreatedAt string `json:"created_at"`
}

type CreateUserOutputV2 struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type createUserOrchestrator struct {
	userRepo  domain.UserRepository
	presenter CreateUserPresenter
}

func NewCrateUserUsecase(
	presenter CreateUserPresenter,
	userRepo domain.UserRepository,
) CreateUserUsecase {
	return &createUserOrchestrator{
		userRepo:  userRepo,
		presenter: presenter,
	}
}

func (uc createUserOrchestrator) CreateUser(context.Context, CreateUserInput) (CreateUserOutput, error) {
	return CreateUserOutput{}, nil
}
