package sql

import (
	"context"
	"time"

	"github.com/achmadAlli/go-simple-boilerplate/adapters/repository"
	"github.com/achmadAlli/go-simple-boilerplate/domain"
	"gorm.io/gorm"
)

const (
	CONTEXT_CREATE_USER = "create_user_transaction"
	CONTEXT_UPDATE_USER = "update_user_transaction"
)

type userModel struct {
	id        string     `gorm:"column:id"`
	name      string     `gorm:"column:name"`
	username  string     `gorm:"column:username"`
	email     string     `gorm:"column:email"`
	password  string     `gorm:"column:password"`
	createdAt *time.Time `gorm:"column:created_at"`
	updatedAt *time.Time `gorm:"column:updated_at"`
}

type UserRepository struct {
	db repository.Transaction
}

func NewUserRepository(db repository.Transaction) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func newUserModel(user domain.User) userModel {
	return userModel{
		id:        string(user.GetID()),
		name:      user.GetName(),
		username:  user.GetUsername(),
		email:     user.GetEmail(),
		password:  user.GetPassword(),
		createdAt: user.GetCreationDate(),
		updatedAt: user.GetUpdateDate(),
	}
}

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
