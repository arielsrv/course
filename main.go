package main

import (
	"course/src/users/application"
	"course/src/users/infrastructure"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", infrastructure.GetUser)
	router.HandleFunc("/ping", PingHandler)
	http.Handle("/", router)

	userRepository := new(infrastructure.HttpUserRepository)
	userService := application.NewUserService(userRepository)
	infrastructure.NewUserController(userService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on :" + port + " ...")
	log.Fatal(server.ListenAndServe())
}

func PingHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode("pong")
}
