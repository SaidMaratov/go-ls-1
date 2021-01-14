package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	nick     string
	room     *room
	commands chan<- command
}

func (c *client) readInput() {

	fmt.Fprintf(c.conn, welcomeIcon+"\n"+"Manual:\n\n"+"/nick - nickname\n"+"/rooms - the list of available rooms\n"+"/join - to create a new room or join the available room\n"+"/quit - leave the server\n\n"+"[ENTER YOUR NAME]:")

	for {
		fmt.Sprint("[" + now.Format("2006-Jan-02 03:04:05") + "][" + c.nick + "]: ")

		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")

		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/nick":
			c.commands <- command{
				id:     CMD_NICK,
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
				args:   args,
			}
		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
				args:   args,
			}
		default:
			if (len(cmd) == 0) || (cmd[0] == '/') {
				c.err(fmt.Errorf("unknown command: %s", cmd))
			} else {
				c.commands <- command{
					id:     CMD_MSG,
					client: c,
					args:   args,
				}
			}
		}

	}

}

func (c *client) err(err error) {
	c.msg(red + "ERR: " + err.Error() + reset)
}

func (c *client) msg(msg string) {
	if len(msg) != 0 {
		c.conn.Write([]byte(msg + "\n"))
	}
	c.conn.Write([]byte("[" + now.Format("2006-Jan-02 03:04:05") + "][" + c.nick + "]:"))
}
