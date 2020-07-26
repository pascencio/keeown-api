package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const socketFile = "/tmp/nginx.socket"
const startedFile = "/tmp/app-initialized"

//Serve ...
func Serve(c context.Context) (e error) {
	if e = os.RemoveAll(socketFile); e != nil {
		log.Error("Error removing socket file")
		return
	}
	if e = os.RemoveAll(startedFile); e != nil {
		log.Error("Error removing app initialized file")
		return
	}
	s := &http.Server{
		Handler: RouteHandler(),
	}
	l, e := net.Listen("unix", socketFile)
	if e != nil {
		log.Error("Error starting server", e)
		return
	}
	defer l.Close()
	go func() {
		s.Serve(l)
	}()
	f, e := os.Create(startedFile)
	if e != nil {
		log.Error("Error starting server", e)
		return
	}
	defer f.Close()
	log.Info("Server started")
	<-c.Done()
	_c, cancel := context.WithTimeout(c, time.Second*5)
	go func() {
		cancel()
	}()
	if e = s.Shutdown(_c); e != nil {
		log.Error("Error stoping server")
		return
	}
	log.Info("Server exited properly")
	if e == http.ErrServerClosed {
		e = nil
	}
	return

}
