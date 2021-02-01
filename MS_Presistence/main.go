package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	pb "github.com/iavealokin/microservices/MS_Generation/user"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
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
	var user User
	conn, err := amqp.Dial("amqp://remote:Cfyz11005310@localhost:5672")
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a channel")
	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("new", true, false, false, false, nil)
	handleError(err, "Couldn't declare `new` queue")
	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Couldn't configure QoS")
	fmt.Println(queue)
	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Couldn't register consumer")
	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID:%d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Recieved a message: %s", d.Body)
			userstruct := User{}
			err := json.Unmarshal(d.Body, &user)
			insertToDB(user)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}
			log.Printf("Result of username %s is password %s", userstruct.Username, userstruct.Password)
			if err := d.Ack(false); err != nil {
				log.Printf("error acknowledging message: %s", err)
			} else {
				log.Printf("Acnowledged messae")
			}
		}
	}()
	<-stopChan
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

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
