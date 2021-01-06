package authorization

import "time"

type User struct {
	Name    string
	History map[time.Time]string
}

func CreateNewAccount(name string) User {
	now := time.Now()
	newHistory := make(map[time.Time]string)

	newHistory[now] = " has joined our chat..."

	newUser := User{
		Name:    name,
		History: newHistory,
	}

	return newUser
}
