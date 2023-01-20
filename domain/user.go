package domain

import (
	"context"
	"time"
)

type UserID string

type User struct {
	id        UserID
	name      string
	username  string
	email     string
	password  string
	createdAt *time.Time
	updatedAt *time.Time
}

type UserRepository interface {
	Store(context.Context, User) User
	Fetch(context.Context) ([]User, error)
	Find(context.Context, User) (User, error)
	Update(context.Context, User) error
	Destroy(context.Context, User) error
}

func NewUser(email string, pass string) User {
	now := time.Now()

	return User{
		id:        "test",
		email:     email,
		password:  pass,
		createdAt: &now,
		updatedAt: &now,
	}
}

func (d *User) GetID() UserID {
	return d.id
}

func (d *User) GetName() string {
	return d.name
}

func (d *User) GetUsername() string {
	return d.username
}

func (d *User) GetEmail() string {
	return d.email
}

func (d *User) SetName(name string) {
	d.name = name
}

func (d *User) SetUsername(username string) {
	d.username = username
}

func (d *User) Setpassword(pass string) {
	d.password = pass
}
