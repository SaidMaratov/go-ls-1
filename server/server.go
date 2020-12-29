package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const welcome = "Welcome to TCP-Chat!"

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
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	fmt.Fprintf(conn, welcome+"\n"+"[ENTER YOUR NAME]:")
	// посмотреть позже
	for scanner.Scan() {
		name := scanner.Text()
		fmt.Fprintf(conn, "Hello, %s\n", name)
		fmt.Fprintf(conn, "%s has joined our chat...\n", name)
		break
	}
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "someone: %s\n", ln)
	}
	defer conn.Close()

	fmt.Println("Code got here.")
}
