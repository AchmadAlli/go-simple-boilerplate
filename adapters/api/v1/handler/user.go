package handler

import (
	"net/http"
	"time"

	"github.com/achmadAlli/go-simple-boilerplate/adapters/api/utils"
	"github.com/achmadAlli/go-simple-boilerplate/adapters/api/v1/request"
	"github.com/achmadAlli/go-simple-boilerplate/usecases/user_usecase"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
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
	var req request.CreateUser
	c, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	user, err := h.createUsecase.CreateUser(c, user_usecase.CreateUserInput{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, utils.NewResponse(user, http.StatusText(http.StatusOK)))
}

func (h UserHandler) Find(ctx echo.Context) error {
	id := ctx.Param("id")

	user, err := h.fetchUsecase.FindUser(ctx.Request().Context(), id)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, utils.NewResponse(user, "ok"))
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
