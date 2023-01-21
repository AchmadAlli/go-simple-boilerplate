package sql

import (
	"time"

	"github.com/achmadAlli/go-simple-boilerplate/adapters/repository"
	"github.com/achmadAlli/go-simple-boilerplate/domain"
)

const (
	CONTEXT_CREATE_USER = "create_user_transaction"
	CONTEXT_UPDATE_USER = "update_user_transaction"
)

type user struct {
	Id        string     `gorm:"column:id"`
	Name      string     `gorm:"column:name"`
	Username  string     `gorm:"column:username"`
	Email     string     `gorm:"column:email"`
	Password  string     `gorm:"column:password"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

type UserRepository struct {
	db repository.Transaction
}

func NewUserRepository(db repository.Transaction) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func newUserModel(d domain.User) user {
	return user{
		Id:        string(d.GetID()),
		Name:      d.GetName(),
		Username:  d.GetUsername(),
		Email:     d.GetEmail(),
		Password:  d.GetPassword(),
		CreatedAt: d.GetCreationDate(),
		UpdatedAt: d.GetUpdateDate(),
	}
}
