package server

import (
	"net"
)

type room struct {
	name    string
	members map[net.Addr]*client
	history string
}

func (r *room) broadcast(sender *client, msg string) {
	for addr, m := range r.members {
		if addr != sender.conn.RemoteAddr() {
			m.msg(msg + "\n")
			m.msg("[" + now.Format("2006-Jan-02 03:04:05") + "][" + m.nick + "]:")
		} else {
			m.msg("[" + now.Format("2006-Jan-02 03:04:05") + "][" + m.nick + "]:")
		}
	}
}
