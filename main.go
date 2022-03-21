package main

import (
	"course/src/users/application"
	"course/src/users/infrastructure"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening...")
	log.Fatal(server.ListenAndServe())
}

func PingHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode("pong")
}
