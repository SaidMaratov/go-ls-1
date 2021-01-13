package main

import (
	"os"

	"./server"
)

func main() {
	var ip string
	args := os.Args[1:]
	if len(args) == 0 {
		ip = ":8989"
	} else {
		ip = args[0]
	}
	server.CreatePort(ip)
}
