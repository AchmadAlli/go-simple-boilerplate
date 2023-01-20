package user_presenters

import (
	"github.com/achmadAlli/go-simple-boilerplate/domain"
	"github.com/achmadAlli/go-simple-boilerplate/usecases/user_usecase"
)

type fetchUserPresenter struct {
}

func NewFetchUserPresenter() user_usecase.FetchUserPresenter {
	return &fetchUserPresenter{}
}

func (p fetchUserPresenter) PresentUser(user domain.User) user_usecase.FetchUserOutput {
	return user_usecase.FetchUserOutput{
		ID:       string(user.GetID()),
		Name:     user.GetName(),
		Username: user.GetUsername(),
		Email:    user.GetEmail(),
	}
}

func (p fetchUserPresenter) PresentUsers(users []domain.User) []user_usecase.FetchUserOutput {
	output := make([]user_usecase.FetchUserOutput, len(users))

	for _, u := range users {
		user := user_usecase.FetchUserOutput{
			ID:       string(u.GetID()),
			Name:     u.GetName(),
			Username: u.GetUsername(),
			Email:    u.GetEmail(),
		}

		output = append(output, user)
	}

	return output
}
