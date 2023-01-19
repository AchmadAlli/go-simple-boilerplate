package usecase

import "github.com/achmadAlli/go-simple-boilerplate/domain/entities"

type User interface {
	Authenticate(auth entities.Auth) error
}
