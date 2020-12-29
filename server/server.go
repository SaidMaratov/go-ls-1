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
		fmt.Fprintf(conn, "%s\n", Users)
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
		user := authorization.CreateNewAccount(name, "has joined our chat...")
		if Users == nil {
			Users = make(map[int]authorization.User)
			i = 0
		} else {
			i = len(Users)
		}
		Users[i] = user
		fmt.Fprintf(conn, "Hello, %s\n", Users[i])
		// fmt.Fprintf(conn, "%s has joined our chat...\n", Users[i])
		fmt.Println(Users)
		break
	}
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Printf("%v - %s\n", Users[i], ln)
		Users[i].History[time.Now()] = ln
		fmt.Fprintf(conn, "%s\n", Users[i].Name)
	}

	defer conn.Close()

	fmt.Println("Code got here.")
}
