package main

import (
	"log"
	"net/http"
	"github.com/gophermasters/x/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gophermasters/bug-free-report/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "10000",
		server: gin.Default(),
	}
}

func (s *Server) RunServer() {
	router := routes.ConfigRoutes(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}

func main(){
	bugsReportServer := NewServer()
	http.HandleFunc("/", handlers.ServeHTML)
	bugsReportServer.RunServer()
}
