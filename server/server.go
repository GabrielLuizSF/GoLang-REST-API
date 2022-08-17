package server

import (
	"log"
	"net/http"
        "user_api-golang/database"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
      _ "github.com/jinzhu/gorm/dialects/mysql"
)


func StartServer(){

      database.InitDB()

      log.Println("Starting the HTTP server on port 5700")

      router := mux.NewRouter().StrictSlash(true)

      initaliseHandlers(router)

      headers := handlers.AllowedHeaders([]string{"X-Request-With","Content-Type","Authorization"})

      methods := handlers.AllowedMethods([]string{"GET","POST","PUT","DELETE"})

      origins := handlers.AllowedOrigins([]string{"*"})

      log.Fatal(http.ListenAndServe(":5700", handlers.CORS(headers,methods,origins)(router)))
}
