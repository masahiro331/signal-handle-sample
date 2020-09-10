package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		done <- true
	}()

	go func() {
		mainLogic(done)
	}()

	<-done
	fmt.Println("finished")
}

func mainLogic(done chan bool) {
	defer func() { done <- true }()
	time.Sleep(time.Second * 3)
}
