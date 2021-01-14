package server

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

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
	if len(args) < 2 || len(args[1]) == 0 {
		c.err(errors.New("Empty name is prohibited!"))
		return
	}
	c.nick = args[1]
	c.msg(fmt.Sprintf(green+"all right, i will call you %s"+reset, c.nick))
}

func (s *server) join(c *client, args []string) {

	if len(args) < 2 || len(args[1]) == 0 {
		c.err(errors.New("Empty name of room is prohibited!"))
		return
	}

	roomName := args[1]

	if len(roomName) > 10 {
		c.err(errors.New("Too much symbols! Try Again."))
		return
	}

	if (c.room != nil) && (c.room.name == roomName) {
		c.msg(fmt.Sprintf(green+"You are already in the %s"+reset, c.room.name))
		return
	}

	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
			history: roomName + ".txt",
		}
		r.createFile(r.history)
		s.rooms[roomName] = r
	}
	r.members[c.conn.RemoteAddr()] = c

	s.quitCurrentRoom(c)
	c.room = r
	if ok {
		s.parseAndReturnHistory(c)
	}
	s.writeToFile(c.room, c.nick+" has joined the room")
	r.broadcast(c, fmt.Sprintf(green+"\n%s has joined the room"+reset, c.nick))
	c.msg(fmt.Sprintf(green+"welcome to %s"+reset, r.name))
}

func (s *server) listRooms(c *client, args []string) {
	if len(s.rooms) == 0 {
		c.msg(green + "There is no room here. Please create a room!" + reset)
		return
	}
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}
	c.msg(fmt.Sprintf(green+"available rooms are: %s"+reset, strings.Join(rooms, ", ")))
}

func (s *server) msg(c *client, args []string) {
	if c.room == nil {
		c.err(errors.New("You must join the room first"))
		return
	}
	if len(args) == 0 {
		c.err(errors.New("Empty messages are prohibited!"))
		return
	}

	s.writeToFile(c.room, "["+now.Format("2006-Jan-02 03:04:05")+"]["+c.nick+"]: "+strings.Join(args, " "))
	c.room.broadcast(c, "\n["+now.Format("2006-Jan-02 03:04:05")+"]["+c.nick+"]: "+strings.Join(args, " "))
	c.msg("")
}

func (s *server) quit(c *client, args []string) {
	log.Printf("client has disconnected: %s", c.conn.RemoteAddr().String())
	s.quitCurrentRoom(c)
	c.msg(green + "You've leaved the server!" + reset)
	c.conn.Close()
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf(green+"\n%s has left the room"+reset, c.nick))
		s.writeToFile(c.room, c.nick+" has left the room")
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

	}
}
