package main

import (
	t "test/types"

	"github.com/bxcodec/faker/v3"
)

func main() {

	var user t.User

	users := *user.CreateUsers(10)

	user.AddNewUser(&users, faker.Email(), faker.FirstName(), faker.LastName())

	user.PrintUsers(&users)

}
