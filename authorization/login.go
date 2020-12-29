package authorization

type User struct {
	// Id   int
	Name string
}

func CreateNewAccount(name string) User {
	// newId := 0
	newUser := User{
		Name: name,
	}

	return newUser
}
