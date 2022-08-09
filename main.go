package main

import (
	"github.com/GabrielLuizSF/GoLang-REST-API/database"
	"github.com/GabrielLuizSF/GoLang-REST-API/server"
)
func main(){
	database.StartDB()
	serverInit := server.NewServer()

	serverInit.RunServer()
	
}