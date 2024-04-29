package types

import (
	"github.com/google/uuid"
)

type User struct {
	Id    string
	Email string
	Profile
}

func NewUser(email string) *User {
	return &User{
		Id:      uuid.New().String(),
		Email:   email,
		Profile: Profile{},
	}
}

func (user *User) GetId() string {
	return user.Id
}

func (user *User) GetEmail() string {
	return user.Email
}

func (user *User) SetId(Id int) *User {
	user.Id = uuid.New().String()
	return user
}

func (user *User) SetEmail(Email string) *User {
	user.Email = Email
	return user
}
