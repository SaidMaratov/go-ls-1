package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	"../authorization"
)

var Users map[int]authorization.User

const welcomeIcon = "" +
	"Welcome to TCP-Chat!\n" +
	"         _nnnn_\n" +
	"        dGGGGMMb\n" +
	"       @p~qp~~qMb\n" +
	"       M|@||@) M|\n" +
	"       @,----.JM|\n" +
	"      JS^\\__/  qKL\n" +
	"     dZP        qKRb\n" +
	"    dZP          qKKb\n" +
	"   fZP            SMMb\n" +
	"   HZM            MMMM\n" +
	"   FqM            MMMM\n" +
	" __| \".        |\\dS\"qML\n" +
	" |    `.       | `' \\Zq\n" +
	"_)      \\.___.,|     .'\n" +
	"\\____   )MMMMMP|   .'\n" +
	"     `-'       `--'\n"

var now = time.Now()

func CreatePort(num string) {

	li, err := net.Listen("tcp", num)
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	fmt.Println("Started server on", num)
	for {
		//var mutex sync.Mutex
		conn, err := li.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}

		go handle(conn)

	}
}

func handle(conn net.Conn) {
	var i int
	scanner := bufio.NewScanner(conn)
	fmt.Fprintf(conn, welcomeIcon+"\n"+"[ENTER YOUR NAME]:")
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
		welcomeMessage := fmt.Sprintf("%v has joined our chat...\n", Users[i].Name)
		fmt.Println(welcomeMessage)
		//channel <- welcomeMessage
		break
	}
	now := time.Now()
	fmt.Fprintf(conn, "[%s][%v]: ", now.Format("2006-Jan-02 03:04:05"), Users[i].Name)

	for scanner.Scan() {
		now = time.Now()

		message := scanner.Text()
		//mes := <-channel
		//fmt.Fprintf(conn, "%s", mes)
		//fmt.Printf("[%s][%v]: %s\n", now.Format("2006-Jan-02 03:04:05"), Users[i].Name, message)
		Users[i].History[time.Now()] = message
		//channel <- fmt.Sprintf("[%v][%v]: %s", now.Format("2006-Jan-02 03:04:05"), Users[i].Name, message)
		//fmt.Print(Users[i].History[time.Now()])
		//channel <- Users[i].History[time.Now()]

		fmt.Fprintf(conn, "[%s][%v]: ", now.Format("2006-Jan-02 03:04:05"), Users[i].Name)
	}

	fmt.Printf("%v has left our chat...\n", Users[i].Name)

	defer conn.Close()

	fmt.Println("Code got here.")
}
