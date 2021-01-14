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

func (s *server) createRoom(roomName string) *room {
	r := &room{
		name:    roomName,
		members: make(map[net.Addr]*client),
		history: roomName + ".txt",
	}
	r.createFile(r.history)
	return r
}
