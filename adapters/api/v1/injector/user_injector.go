package injector

import (
	"github.com/achmadAlli/go-simple-boilerplate/adapters/api/v1/handler"
	"github.com/achmadAlli/go-simple-boilerplate/adapters/presenter/user_presenters"
	"github.com/achmadAlli/go-simple-boilerplate/usecases/user_usecase"
)

func ProvideUserHandler() *handler.UserHandler {
	fetchPresenter := user_presenters.NewFetchUserPresenter()
	creationPresenter := user_presenters.NewCreateUserPresenter()

	createUserUsecase := user_usecase.NewCrateUserUsecase(creationPresenter)
	fetchUserUsecase := user_usecase.NewFetchUserUsecase(fetchPresenter)

	return handler.NewUserHandler(
		createUserUsecase,
		fetchUserUsecase,
	)
}
