package authorization

import "time"

type User struct {
	Name    string
	Id      int
	History map[time.Time]string
}

func CreateNewAccount(name string, id int) User {
	now := time.Now()
	newHistory := make(map[time.Time]string)

	newHistory[now] = " has joined our chat..."

	newUser := User{
		Name:    name,
		Id:      id,
		History: newHistory,
	}

	return newUser
}
