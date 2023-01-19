package repository

import (
	"context"

	"github.com/achmadAlli/go-simple-boilerplate/domain/entities"
)

type UserRepository interface {
	Store(context.Context, entities.User) entities.User
	Fetch(context.Context) ([]entities.User, error)
	Find(context.Context, entities.User) (entities.User, error)
	Update(context.Context, entities.User) error
	Delete(context.Context, entities.User) error
}
