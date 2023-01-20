package handler

import (
	"net/http"

	"github.com/achmadAlli/go-simple-boilerplate/adapters/api/utils"
	"github.com/achmadAlli/go-simple-boilerplate/domain"
	"github.com/achmadAlli/go-simple-boilerplate/usecases/user_usecase"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	createUsecase user_usecase.CreateUserUsecase
	fetchUsecase  user_usecase.FetchUserUsecase
}

func NewUserHandler(createUsecase user_usecase.CreateUserUsecase, fetchUsecase user_usecase.FetchUserUsecase) *UserHandler {
	return &UserHandler{
		createUsecase: createUsecase,
		fetchUsecase:  fetchUsecase,
	}
}

func (h UserHandler) Store(ctx echo.Context) error {
	user := domain.User{}
	return ctx.JSON(http.StatusOK, utils.NewResponse(user, http.StatusText(http.StatusOK)))
}

func (h UserHandler) Find(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, utils.NewResponse(nil, "ok"))
}
func (h UserHandler) Fetch(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, utils.NewResponse(nil, "ok"))

}
func (h UserHandler) Update(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, utils.NewResponse(nil, "ok"))

}
func (h UserHandler) Destroy(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, utils.NewResponse(nil, "ok"))

}
