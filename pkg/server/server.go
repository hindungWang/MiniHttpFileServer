package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Path string
	Port string
}

func (s *Server) Start() error {
	http.Handle("/", http.FileServer(http.Dir(s.Path)))
	log.Infof("server starting on %s,exposed path at %s...", s.Port, s.Path)
	err := http.ListenAndServe(fmt.Sprint(":", s.Port), nil)
	if err != nil {
		log.Error(err)
	}
	return err
}
