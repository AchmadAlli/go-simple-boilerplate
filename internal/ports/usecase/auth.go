package usecase

import "github.com/achmadAlli/go-simple-boilerplate/internal/entities"

type User interface {
	Authenticate(auth entities.Auth) error
}
