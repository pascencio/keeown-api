package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/pascencio/keeown-api/server"
	log "github.com/sirupsen/logrus"
)

func setupLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		ForceColors:            true,
		DisableLevelTruncation: true,
		QuoteEmptyFields:       true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
}

func main() {
	setupLogger()
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	c, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-s
		log.Info("System call", oscall)
		cancel()
	}()
	if e := server.Serve(c); e != nil {
		log.Error("Error starting server")
	}
}
