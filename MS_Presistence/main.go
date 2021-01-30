package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"

	pb "github.com/iavealokin/microservices/MS_Generation/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// User struct ...
type User struct {
	Login    string `json:"login"`
	Username string `json:"username"`
	Surname  string `json:"surname"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
}

//Структура нашего gRPC сервера
type server struct {
}

const port = ":20100"

func (s *server) SendPass(ctx context.Context, in *pb.MsgRequest) (*pb.MsgReply, error) {
	var user User
	_ = json.Unmarshal([]byte(in.Message), &user)
	fmt.Println(user)
	insertToDB(user)
	return &pb.MsgReply{Sent: true}, nil
}

func main() {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen", err)
	}
	log.Printf("start listening for emails at port %s", port)

	rpcserv := grpc.NewServer()

	//Регистрируем связку сервер + listener
	pb.RegisterUserServer(rpcserv, &server{})
	reflection.Register(rpcserv)

	//Запускаемся и ждём RPC-запросы
	err = rpcserv.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve", err)
	}
}

func insertToDB(user User) {
	db, err := sql.Open("postgres", "postgres://remote:Cfyz11005310@localhost/microservices")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(fmt.Sprintf("INSERT INTO users(login,name,surname,birthday,password) values('%s','%s','%s','%s','%s')",
		user.Login, user.Username, user.Surname, user.Birthday, user.Password))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}
