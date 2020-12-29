package authorization

import "time"

type User struct {
	// Id   int
	Name string
	// History map[string]string
	History map[time.Time]string
}

func CreateNewAccount(name string) User {
	now := time.Now()
	// newHistory := make(map[string]string)
	newHistory := make(map[time.Time]string)

	// newHistory[now.Format("2006-Jan-02 03:04:05")] = " has joined our chat..."
	newHistory[now] = " has joined our chat..."

	newUser := User{
		Name:    name,
		History: newHistory,
	}

	return newUser
}
