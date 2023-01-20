package user_presenters

import "github.com/achmadAlli/go-simple-boilerplate/usecases/user_usecase"

type createUserPresenter struct {
}

func NewCreateUserPresenter() user_usecase.CreateUserPresenter {
	return &createUserPresenter{}
}

func (p createUserPresenter) PresentCreatedUser() user_usecase.CreateUserOutput {
	return user_usecase.CreateUserOutput{}
}

func (p createUserPresenter) PresentCreatedUserV2() user_usecase.CreateUserOutputV2 {
	return user_usecase.CreateUserOutputV2{}
}
