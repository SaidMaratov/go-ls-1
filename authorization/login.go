package authorization

import "time"

type User struct {
	// Id   int
	Name    string
	History map[time.Time]string
}

func CreateNewAccount(name string, message string) User {

	newHistory := make(map[time.Time]string)
	newHistory[time.Now()] = message

	newUser := User{
		Name:    name,
		History: newHistory,
	}

	return newUser
}
