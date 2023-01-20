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
		id:        UserID(NewUUID()),
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
	now := time.Now()
	d.name = name
	d.createdAt = &now
}

func (d *User) SetUsername(username string) {
	now := time.Now()
	d.username = username
	d.createdAt = &now
}

func (d *User) SetPassword(pass string) {
	now := time.Now()
	d.password = pass
	d.createdAt = &now
}
