package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	"../authorization"
)

const welcome = "Welcome to TCP-Chat!"

var Users map[int]authorization.User

func CreatePort(num string) {
	li, err := net.Listen("tcp", num)
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	fmt.Println("Server is Listening", num)
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		if conn != nil {
			fmt.Println("coming")
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	var i int
	scanner := bufio.NewScanner(conn)
	fmt.Fprintf(conn, welcome+"\n"+"[ENTER YOUR NAME]:")
	// посмотреть позже
	for scanner.Scan() {

		name := scanner.Text()
		user := authorization.CreateNewAccount(name)
		if Users == nil {
			Users = make(map[int]authorization.User)
			i = 0
		} else {
			i = len(Users)
		}
		Users[i] = user
		// fmt.Fprintf(conn, "Hello, %s\n", Users[i].Name)
		fmt.Println(Users)
		break
	}
	now := time.Now()
	fmt.Fprintf(conn, "[%s][%v]: ", now.Format("2006-Jan-02 03:04:05"), Users[i].Name)
	for scanner.Scan() {
		now = time.Now()

		ln := scanner.Text()

		fmt.Printf("[%s][%v]: %s\n", now.Format("2006-Jan-02 03:04:05"), Users[i].Name, ln)
		Users[i].History[now.Format("2006-Jan-02 03:04:05")] = ln
		// fmt.Fprintf(conn, "[%s][%v]: %s\n", now.Format("2006-Jan-02 03:04:05"), Users[i].Name, ln)
		fmt.Fprintf(conn, "[%s][%v]: ", now.Format("2006-Jan-02 03:04:05"), Users[i].Name)
	}

	defer conn.Close()

	fmt.Println("Code got here.")
}
