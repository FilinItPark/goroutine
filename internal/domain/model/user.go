package model

import "fmt"

type User struct {
	id   int64
	Name string
}

func NewUser(id int64, name string) User {
	user := User{
		id:   id,
		Name: name,
	}

	fmt.Printf("%p", (&user))
	fmt.Println()
	return user
}
