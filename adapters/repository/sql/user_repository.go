package sql

import (
	"context"

	"github.com/achmadAlli/go-simple-boilerplate/domain"
	"gorm.io/gorm"
)

func (r *UserRepository) Store(ctx context.Context, user domain.User) (domain.User, error) {
	payload := newUserModel(user)
	err := r.db.FromContext(ctx, CONTEXT_CREATE_USER).(*gorm.DB).Create(&payload).Error

	return user, err
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
