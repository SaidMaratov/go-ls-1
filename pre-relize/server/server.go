package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"../authorization"
)

var Users map[int]authorization.User

var now = time.Now()

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("new client has connected: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.commands,
	}
	c.readInput()
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)
		}
	}
}
func (s *server) nick(c *client, args []string) {
	c.nick = args[1]
	c.msg(fmt.Sprintf("all right, i will call you %s", c.nick))
}

func (s *server) join(c *client, args []string) {
	roomName := args[1]
	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}
	r.members[c.conn.RemoteAddr()] = c
	s.quitCurrentRoom(c)
	c.room = r
	r.broadcast(c, fmt.Sprintf("%s has joined the room", c.nick))
	c.msg(fmt.Sprintf("welcome to %s", r.name))
}

func (s *server) listRooms(c *client, args []string) {
	if len(s.rooms) == 0 {
		c.msg("There is no room here. Please create a room!")
		return
	}
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}
	c.msg(fmt.Sprintf("available rooms are: %s", strings.Join(rooms, ", ")))
}

func (s *server) msg(c *client, args []string) {
	if c.room == nil {
		c.err(errors.New("You must join the room first"))
		return
	}
	c.room.broadcast(c, c.nick+": "+strings.Join(args, " "))
}

func (s *server) quit(c *client, args []string) {
	log.Printf("client has disconnected: %s", c.conn.RemoteAddr().String())
	s.quitCurrentRoom(c)
	c.msg("You've leaved the server!")
	c.conn.Close()
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s has left the room", c.nick))
		c.room = &room{}
	}
}

func CreatePort(num string) {

	s := newServer()
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

		//go handle(conn)

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
