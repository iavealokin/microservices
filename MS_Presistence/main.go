package main

import (
	"log"
	"net"

	ps "MS_Generation/UserService/proto/"

	"google.golang.org/grpc"
)

type User struct {
	Login    string `json:"login"`
	Username string `json:"username"`
	Surname  string `json:"surname"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
}

func main() {
	server := grpc.NewServer()
	instance := new(User)
	ps.RegisterPasswordGeneratorServiceServer(server, instance)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Unable to create grpc listener:", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
