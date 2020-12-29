package authorization

import "time"

type User struct {
	// Id   int
	Name    string
	History map[string]string
}

func CreateNewAccount(name string) User {
	now := time.Now()
	newHistory := make(map[string]string)

	newHistory[now.Format("2006-Jan-02 03:04:05")] = " has joined our chat..."

	newUser := User{
		Name:    name,
		History: newHistory,
	}

	return newUser
}

// func CreateNewMessage(date string, message string) User {
// 	now := time.Now()
// 	newHistory := make(map[string]string)

// 	newHistory[now.Format("2006-Jan-02 03:04:05")] = " has joined our chat..."

// 	newUser := User{
// 		Name:    name,
// 		History: newHistory,
// 	}

// 	return newUser
// }
