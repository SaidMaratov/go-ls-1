package main

import (
	"os"

	"./server"
)

func main() {
	args := os.Args[1:]
	server.CreatePort(args[0])
}
