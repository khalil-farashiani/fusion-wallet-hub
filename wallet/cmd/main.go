package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sig := waitExitSignal()
	log.Println(sig.String())
}

// waitExitSignal get os signals
func waitExitSignal() os.Signal {
	quit := make(chan os.Signal, 6)
	signal.Notify(quit, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	return <-quit
}
