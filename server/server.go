package server

import (
"log"
"github.com/gin-gonic/gin"
"github.com/GabrielLuizSF/GoLang-REST-API/server/routes"

)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server{
	return Server{
		port: "8500",
		server: gin.Default(),
	}
}
func (serverInit *Server) Run(){
	router := routes.ConfigRoutes(serverInit.server)

	log.Printf("Server runing at port :  %v",serverInit.port)
    log.Fatal(router.Run(":"+serverInit.port))

}