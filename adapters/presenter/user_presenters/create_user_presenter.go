package user_presenters

import (
	"github.com/achmadAlli/go-simple-boilerplate/domain"
	"github.com/achmadAlli/go-simple-boilerplate/usecases/user_usecase"
)

type createUserPresenter struct {
}

func NewCreateUserPresenter() user_usecase.CreateUserPresenter {
	return &createUserPresenter{}
}

func (p createUserPresenter) PresentCreatedUser(user domain.User) user_usecase.CreateUserOutput {
	var creation string

	if user.GetCreationDate() != nil {
		creation = user.GetCreationDate().Format(domain.DATE_TIME_FORMAT_ISO_DATE)
	}

	return user_usecase.CreateUserOutput{
		ID:        string(user.GetID()),
		Name:      user.GetName(),
		Username:  user.GetUsername(),
		Email:     user.GetEmail(),
		CreatedAt: creation,
	}
}

func (p createUserPresenter) PresentCreatedUserV2(user domain.User) user_usecase.CreateUserOutputV2 {
	var (
		creation string
		update   string
	)

	if user.GetCreationDate() != nil {
		creation = user.GetCreationDate().Format(domain.DATE_TIME_FORMAT_ISO_DATE)
	}

	if user.GetUpdateDate() != nil {
		creation = user.GetUpdateDate().Format(domain.DATE_TIME_FORMAT_ISO_DATE)
	}

	return user_usecase.CreateUserOutputV2{
		ID:        string(user.GetID()),
		Name:      user.GetName(),
		Username:  user.GetUsername(),
		Email:     user.GetEmail(),
		CreatedAt: creation,
		UpdatedAt: update,
	}
}
