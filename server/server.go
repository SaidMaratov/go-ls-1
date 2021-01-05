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

type server struct {
	rooms map[string]*room
	commands chan command
}

func newServer()*server{
	return &server{
		rooms: make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) newClient(conn net.Conn){
	log.Printf("new client has connected: %s", conn.RemoteAddr().String())
	c:=&client{
		conn: conn,
		nick: "anonymous",
		commands: s.commands,
	}
	c.readInput()
}

func (s *server) run(){
	for cmd := range s.commands{
		switch cmd.id{
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_Quit:
			s.quit(cmd.client, cmd.args)
		}
	}
}



func CreatePort(num string) {

	s:= newServer()
	go s.run()

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
		go s.newClient(conn)

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
