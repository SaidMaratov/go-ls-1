package server

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"../authorization"
)

const welcome = "Welcome to TCP-Chat!"

var i = 0

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

	scanner := bufio.NewScanner(conn)
	fmt.Fprintf(conn, welcome+"\n"+"[ENTER YOUR NAME]:")
	// посмотреть позже
	for scanner.Scan() {
		name := scanner.Text()
		user := authorization.CreateNewAccount(name)
		if Users == nil {
			Users = make(map[int]authorization.User)
		}
		Users[i] = user
		fmt.Fprintf(conn, "Hello, %s\n", Users[i])
		fmt.Fprintf(conn, "%s has joined our chat...\n", Users[i])
		fmt.Println(Users)
		i++
		break
	}
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "%s\n", Users[i])
	}
	defer conn.Close()

	fmt.Println("Code got here.")
}
