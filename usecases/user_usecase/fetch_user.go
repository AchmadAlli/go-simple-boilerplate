package user_usecase

import (
	"context"

	"github.com/achmadAlli/go-simple-boilerplate/domain"
)

type FetchUserUsecase interface {
	FindUser(ctx context.Context, ID string) (FetchUserOutput, error)
	FetchUsers(ctx context.Context, filter FetchUserInput) ([]FetchUserOutput, error)
}

type FetchUserPresenter interface {
	PresentUser(user domain.User) FetchUserOutput
	PresentUsers(users []domain.User) []FetchUserOutput
}

type FetchUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FetchUserOutput struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type fetchUserOrchestrator struct {
	userRepo  interface{}
	presenter FetchUserPresenter
}

func NewFetchUserUsecase(presenter FetchUserPresenter) FetchUserUsecase {
	return &fetchUserOrchestrator{
		userRepo:  nil,
		presenter: presenter,
	}
}

func (uc fetchUserOrchestrator) FindUser(ctx context.Context, ID string) (FetchUserOutput, error) {
	return FetchUserOutput{}, nil
}

func (uc fetchUserOrchestrator) FetchUsers(ctx context.Context, filter FetchUserInput) ([]FetchUserOutput, error) {
	return []FetchUserOutput{}, nil
}
