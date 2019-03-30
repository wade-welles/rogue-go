package rogue

import (
	"os"
	"os/signal"
	"syscall"
)

func NewSignalHandler() {
	shutdownSignals := make(chan bool, 1)
	go signalHandler(shutdownSignals)
}

func signalHandler(shutdownSignals <-chan bool) {
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT, os.Kill)
	for {
		s := <-signals
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, os.Kill:
			PrintLog("SHUTDOWN", true, White("User initiated shutdown by pressing CTRL+C..."))
			os.Exit(0)
		}
	}
}
