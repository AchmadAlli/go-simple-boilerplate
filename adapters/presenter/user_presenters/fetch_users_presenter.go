package user_presenters

import "github.com/achmadAlli/go-simple-boilerplate/usecases/user_usecase"

type fetchUserPresenter struct {
}

func NewFetchUserPresenter() user_usecase.FetchUserPresenter {
	return &fetchUserPresenter{}
}

func (p fetchUserPresenter) PresentUser() user_usecase.FetchUserOutput {
	return user_usecase.FetchUserOutput{}
}

func (p fetchUserPresenter) PresentUsers() []user_usecase.FetchUserOutput {
	return []user_usecase.FetchUserOutput{}
}
