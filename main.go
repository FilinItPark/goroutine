package main

type User struct {
	ID   int64
	Name string
}

func (u *User) getName(other string) (string, error) {
	return u.Name + " " + other, nil
}

func main() {

}
