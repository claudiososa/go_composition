package types

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
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

func (user *User) PrintUsers(users *map[string]User) {
	fmt.Println("----------------------------------------------------------------------------------------------------")
	fmt.Printf("%-42s %-25s %-20s %-20s\n", "Id", "Email", "Firstname", "Lastname")
	fmt.Println("----------------------------------------------------------------------------------------------------")

	for item := range *users {
		user := (*users)[item]
		fmt.Printf("%-42s %-25s %-20s %-20s\n", user.Id, user.Email, user.Profile.GetName(), user.Profile.GetSurname())
	}
}

func (user *User) AddNewUser(users *map[string]User, email string, firstname string, lastname string) {
	newUser := NewUser(email)
	newUser.Profile.SetName(firstname)
	newUser.Profile.SetSurname(lastname)
	(*users)[newUser.Id] = *newUser
}

func (user *User) CreateUsers(quantity int) *map[string]User {
	users := make(map[string]User)

	for i := 0; i < quantity; i++ {

		newUser := NewUser(faker.Email())

		newUser.Profile.SetName(faker.FirstName())
		newUser.Profile.SetSurname(faker.LastName())

		users[newUser.Id] = *newUser
	}
	return &users
}
