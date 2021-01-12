package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"../authorization"
)

var Users map[int]authorization.User
var History map[time.Time]string

func CreatePort(port string) {
	li, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	fmt.Println("Server is Listening", port)

	runServer(li)
}

func runServer(li net.Listener) {

	//We read channels for incoming connections, dead connections and messages
	aconns := make(map[net.Conn]int)
	conns := make(chan net.Conn)
	dconns := make(chan net.Conn)
	msgs := make(chan string)
	i := 1

	go func() { //Create connections, when somebody call
		for {
			conn, err := li.Accept()
			if err != nil {
				log.Println(err.Error())
			}
			conns <- conn
		}
	}()

	Users = make(map[int]authorization.User)

	for {
		select {
		//Read incoming connections
		case conn := <-conns:
			aconns[conn] = i
			fmt.Fprintln(conn, welcomeIcon)
			fmt.Fprint(conn, "[ENTER YOUR NAME]: ")
			// once we have the connection, we start reading messages from it
			go func(conn net.Conn, i int) {
				rd := bufio.NewReader(conn)
				for {
					message, err := rd.ReadString('\n')
					if err != nil {
						break
					}
					_, ok := Users[i] // Checking for a user with the id = i, if there is now user with the given id, ok = false
					if !ok {
						name := message[0 : len(message)-1] // The first message would be the name
						user := authorization.CreateNewAccount(name, i)
						Users[i] = user
						welcomeMessage := fmt.Sprintf("%v has joined our chat...\n", Users[i].Name)
						msgs <- fmt.Sprintf("%v%s", i, welcomeMessage)
						now := time.Now()
						fmt.Fprintf(conn, "[%s][%v]: ", now.Format("2006-Jan-02 03:04:05"), Users[i].Name)
					} else {
						Users[i].History[time.Now()] = message
						now := time.Now()
						msgs <- fmt.Sprintf("%v[%s][%v]: %s", i, now.Format("2006-Jan-02 03:04:05"), Users[i].Name, message)
						fmt.Fprintf(conn, "[%s][%v]: ", now.Format("2006-Jan-02 03:04:05"), Users[i].Name)
					}
				}
				// Done reading from it
				dconns <- conn
			}(conn, i)
			i++
		case msg := <-msgs:
			// we have to broadcast it to all connections
			for conn, i := range aconns {
				num, _ := strconv.Atoi(string(msg[0]))
				if i != num {
					conn.Write([]byte("\n"))
					conn.Write([]byte(msg[1:]))
					now := time.Now()
					text := "[" + now.Format("2006-Jan-02 03:04:05") + "][" + Users[i].Name + "]: "
					conn.Write([]byte(text))
				}
			}
		case dconn := <-dconns:
			num := aconns[dconn]
			log.Printf("%s has left our chat...\n", Users[num].Name)
			// we have to broadcast the msg that client is gone to all connections
			for conn := range aconns {
				fmt.Fprintf(conn, "\n%s has left our chat...\n", Users[num].Name)
				now := time.Now()
				text := "[" + now.Format("2006-Jan-02 03:04:05") + "][" + Users[aconns[conn]].Name + "]: "
				conn.Write([]byte(text))
			}
			delete(aconns, dconn)
		}
	}

}

func handle() {

}
