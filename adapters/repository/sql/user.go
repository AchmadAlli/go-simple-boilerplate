package sql

import (
	"context"

	"github.com/achmadAlli/go-simple-boilerplate/adapters/repository"
	"github.com/achmadAlli/go-simple-boilerplate/domain"
)

const (
	USER_CONTEXT = "user_transaction"
)

type UserRepository struct {
	db repository.Transaction
}

func NewUserRepository(db repository.Transaction) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Store(ctx context.Context, User domain.User) domain.User {
	return domain.User{}
}
func (r *UserRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	return []domain.User{}, nil
}
func (r *UserRepository) Find(ctx context.Context, user domain.User) (domain.User, error) {
	return domain.User{}, nil
}
func (r *UserRepository) Update(context.Context, domain.User) error {
	return nil
}
func (r *UserRepository) Destroy(context.Context, domain.User) error {
	return nil
}
