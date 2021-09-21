package main

import (
	"botmanager/server"
	"fmt"
	"github.com/akhilesh18992/glog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Hey")
	s := server.NewServer()
	go func() {
		err := s.StartServer()
		if err != nil {

		}
	}()
	waitForTermination(s)
}

func waitForTermination(s *server.Server) {
	// Starting server
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGABRT, syscall.SIGSEGV)
	exit := make(chan bool, 1)

	go func() {
		glog.V(1).Info("shutting down with Signal:", <-c)
		glog.Flush()
		s.StopServer()
		exit <- true
	}()
	<-exit
}
