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
			m.msg(msg)
		}
	}
}
