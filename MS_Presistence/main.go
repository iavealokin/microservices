package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/iavealokin/microservices/MS_Generation/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// User struct ...
type User struct {
	Login    string
	Username string
	Surname  string
	Birthday string
	Password string
}

//Структура нашего gRPC сервера
type server struct {
}

const port = ":20100"

/*

Методы структуры SendPass и RetrievePass принимают контекст и входящее сообщение,
формат которого мы описали в proto-файле вот так:

message MsgRequest {
    string to = 1;
    string code = 2;
}

Формат ответного сообщения в прото-файле был такой:

message MsgReply {
    bool sent = 1;
}

*/
func (s *server) SendPass(ctx context.Context, in *pb.MsgRequest) (*pb.MsgReply, error) {

	fmt.Println(in)

	return &pb.MsgReply{Sent: true}, nil
}

func (s *server) RetrievePass(ctx context.Context, in *pb.MsgRequest) (*pb.MsgReply, error) {

	//А здесь ответим false

	return &pb.MsgReply{Sent: false}, nil
}

func main() {

	//Указываем на каком порту будем слушать запросы
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen", err)
	}
	log.Printf("start listening for emails at port %s", port)

	//Создаём новый grpc сервер
	rpcserv := grpc.NewServer()

	//Регистрируем связку сервер + listener
	pb.RegisterUserSenderServer(rpcserv, &server{})
	reflection.Register(rpcserv)

	//Запускаемся и ждём RPC-запросы
	err = rpcserv.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve", err)
	}
}
