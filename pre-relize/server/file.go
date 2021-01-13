package server

import (
	"fmt"
	"log"
	"os"
	"time"
)

func (r *room) createFile(name string) {
	_, err := os.Create(name)
	if err != nil {
		fmt.Println(err.Error())
	}

	historyFile, err := os.OpenFile(r.history, os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		log.Println(err.Error())
	}

	now := time.Now()
	_, err = historyFile.WriteString("Room is created: " + now.Format("2006-Jan-02 03:04:05"))
	if err != nil {
		log.Println(err.Error())
	}
	historyFile.Close()
	log.Printf("Room file is created: %s\n", name)
}

func (s *server) writeToFile(r *room, msg string) {
	historyFile, err := os.OpenFile(r.history, os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		log.Println(err.Error())
	}

	_, err = historyFile.WriteString("\n" + msg)
	if err != nil {
		log.Println(err.Error())
	}
	historyFile.Close()
}
