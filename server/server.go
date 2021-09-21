package server

import (
	"botmanager/httphandlers"
	"fmt"
	"github.com/akhilesh18992/glog"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer() error {
	glog.V(1).Infof("Started rabbit server")

	router := httprouter.New()
	router.POST("/set", httphandlers.HelloTestSet)
	router.GET("/get/{key}", httphandlers.HelloTestGet)
	router.PanicHandler = PanicHandler
	return http.ListenAndServe(":8080", router)
}

func (s *Server) StopServer() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
}

func PanicHandler(w http.ResponseWriter, r *http.Request, p interface{}) {
	glog.Errorf("Panic occurred: %s", p)
}