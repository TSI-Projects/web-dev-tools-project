package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/AndrejsPon00/web-dev-tools/backend/selenium"
	"github.com/AndrejsPon00/web-dev-tools/backend/server"
)

func main() {
	closeChan := make(chan os.Signal, 1)
	signal.Notify(closeChan, os.Interrupt)

	go func() {
		<-closeChan
		log.Println("Shutting down")
		close(closeChan)
	}()

	go selenium.StartServer(closeChan)
	go server.Start()

	<-closeChan
}
