package user

import "github.com/achmadAlli/go-simple-boilerplate/domain/ports/repository"

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase() *userUsecase {
	return &userUsecase{
		userRepo: nil,
	}
}
