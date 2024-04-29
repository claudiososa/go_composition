package main

import (
	"fmt"
	t "test/types"

	"github.com/bxcodec/faker/v3"
)

func createUsers(quantity int) *map[string]t.User {
	users := make(map[string]t.User)

	for i := 0; i < quantity; i++ {

		newUser := t.NewUser(faker.Email())

		newUser.Profile.SetName(faker.FirstName())
		newUser.Profile.SetSurname(faker.LastName())

		users[newUser.Id] = *newUser
	}
	return &users
}

func addNewUser(users *map[string]t.User, email string, firstname string, lastname string) {
	newUser := t.NewUser(email)
	newUser.Profile.SetName(firstname)
	newUser.Profile.SetSurname(lastname)
	(*users)[newUser.Id] = *newUser
}

func printUsers(users *map[string]t.User) {
	fmt.Println("----------------------------------------------------------------------------------------------------")
	fmt.Printf("%-42s %-25s %-20s %-20s\n", "Id", "Email", "Firstname", "Lastname")
	fmt.Println("----------------------------------------------------------------------------------------------------")

	for item := range *users {
		user := (*users)[item]
		fmt.Printf("%-42s %-25s %-20s %-20s\n", user.Id, user.Email, user.Profile.GetName(), user.Profile.GetSurname())
	}
}

func main() {

	users := *createUsers(10)

	addNewUser(&users, faker.Email(), faker.FirstName(), faker.LastName())

	printUsers(&users)

}
