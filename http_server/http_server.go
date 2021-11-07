package http_server

import (
	"context"
	"log"
	"net/http"
	"time"
)
import (
	"github.com/gin-gonic/gin"
)

type SettingFunc func(server *Server)

type Server struct {
	httpServer *http.Server
	gin        *gin.Engine

	addr  string
	isTls bool
}

func SetAddress(addr string) SettingFunc {
	return func(server *Server) {
		server.addr = addr
	}
}

func Routing(r Router) SettingFunc {
	return func(server *Server) {
		r.GinRouting(server.gin)
	}
}

func SetTls() SettingFunc {
	return func(server *Server) {
		server.isTls = true
	}
}

func NewServer(settings ...SettingFunc) *Server {
	server := &Server{
		gin: gin.New(),
	}

	for _, setting := range settings {
		setting(server)
	}

	server.httpServer = &http.Server{
		Addr:      server.addr,
		Handler:   server.gin,
		TLSConfig: nil,
	}
	return server
}

func (s *Server) Run() error {
	var err error = nil
	if s.isTls {
		panic("not impl")
	} else {
		err = s.httpServer.ListenAndServe()
	}
	return err
}

func (s *Server) Stop(c context.Context) {
	ctx, cancel := context.WithTimeout(c, 3*time.Second)
	defer cancel()
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		log.Fatal("err :", err)
	}
	log.Println("http server stop")
}
