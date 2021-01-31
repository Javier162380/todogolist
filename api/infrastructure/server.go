package api

import (
	"fmt"
	"todogolist/repository"

	"github.com/gin-gonic/gin"
)

//Server Main infrastructure server of the API.
type Server struct {
	engine      *gin.Engine
	httpAddress string
	repository  repository.Repository
}

//NewServer create a new service instance of the api.
func NewServer(host string, port uint, repository repository.Repository) Server {

	srv := Server{
		engine:      gin.New(),
		httpAddress: fmt.Sprintf("%s:%d", host, port),
		repository:  repository,
	}

	srv.registerRoutes()

	return srv

}

func (s *Server) registerRoutes() {
	s.engine.GET("health/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World"})
	})

}

//Starts a new gin server.
func (s *Server) Run() {
	s.engine.Run()
}
