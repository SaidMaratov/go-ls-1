package authorization

import "time"

type User struct {
	// Id   int
	Name    string
	History map[string]string
}

func CreateNewAccount(name string, message string) User {
	now := time.Now()
	newHistory := make(map[string]string)
	newHistory[now.Format("2006-Jan-02 Monday 03:04:05")] = message

	newUser := User{
		Name:    name,
		History: newHistory,
	}

	return newUser
}
