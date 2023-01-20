package injector

import (
	"github.com/achmadAlli/go-simple-boilerplate/adapters/api/v1/handler"
	"github.com/achmadAlli/go-simple-boilerplate/adapters/presenter/user_presenters"
	"github.com/achmadAlli/go-simple-boilerplate/adapters/repository/sql"
	"github.com/achmadAlli/go-simple-boilerplate/config"
	"github.com/achmadAlli/go-simple-boilerplate/infrastructures/database"
	"github.com/achmadAlli/go-simple-boilerplate/usecases/user_usecase"
)

func ProvideUserHandler() *handler.UserHandler {
	fetchPresenter := user_presenters.NewFetchUserPresenter()
	creationPresenter := user_presenters.NewCreateUserPresenter()
	userRepo := sql.NewUserRepository(database.NewTransaction(config.DBConfig.GetUserDB()))

	createUserUsecase := user_usecase.NewCrateUserUsecase(creationPresenter, userRepo)
	fetchUserUsecase := user_usecase.NewFetchUserUsecase(fetchPresenter, userRepo)

	return handler.NewUserHandler(
		createUserUsecase,
		fetchUserUsecase,
	)
}
