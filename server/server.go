package server

import (
	"apiGolang/server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

//Struct do Server
//utiliza o gin.Engine para criar o server
type Server struct {
	port   string
	server *gin.Engine
}

//Constructor
func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {

	router := routes.ConfigRoutes(s.server)

	log.Println("Server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))

}
