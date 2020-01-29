package util


import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

// WaitInterruption helps to wait OS interruption more easier way when custom signals handling is not required
func WaitInterruption(cb func()) {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			log.Warn("Received an interrupt, stopping...")
			cb()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}