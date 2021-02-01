package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

//User struct ...
type User struct {
	Login    string `json:"login"`
	Username string `json:"username"`
	Surname  string `json:"surname"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
}

var (
	lowerCharSet       = "abcdedfghijklmnopqrst"
	upperCharSet       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet     = "!@#$%&*"
	numberSet          = "0123456789"
	allCharSet         = lowerCharSet + upperCharSet + specialCharSet + numberSet
	delay          int = 20
)

func main() {

	go consumerAmqp()

	for {
		log.Printf("THe delay in infinity lop as : %d", delay)
		time.Sleep(time.Duration(delay) * time.Second)
		go getUser()
	}
}

func consumerAmqp() {
	conn, err := amqp.Dial("amqp://remote:Cfyz11005310@localhost:5672")
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a channel")
	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("delay", true, false, false, false, nil)
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
	var delaymap = make(map[string]int)
	go func() {
		log.Printf("Consumer ready, PID:%d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Recieved a message: %s", d.Body)
			err := json.Unmarshal(d.Body, &delaymap)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}
			log.Printf("Delay is %d", delaymap["delay"])
			delay = delaymap["delay"]
			log.Printf("this is the delay : %d", delay)
			if err := d.Ack(false); err != nil {
				log.Printf("error acknowledging message: %s", err)
			} else {
				log.Printf("Acnowledged message")
			}
		}

	}()
	<-stopChan
}
func getUser() {
	fmt.Println("getUser() is started")
	var usernames = map[int]string{
		0:  "John",
		1:  "David",
		2:  "Steve",
		3:  "Peter",
		4:  "Elena",
		5:  "Caroline",
		6:  "Vanda",
		7:  "William",
		8:  "Harper",
		9:  "Evelyn",
		10: "Jackson",
		11: "Jack",
		12: "James",
		13: "Mason",
		14: "Ella",
		15: "Avery",
	}
	var surnames = map[int]string{
		0:  "Davidson",
		1:  "Maksimoff",
		2:  "Rodgers",
		3:  "Stark",
		4:  "Quill",
		5:  "Doe",
		6:  "Jobs",
		7:  "Aaron",
		8:  "Abbot",
		9:  "Lang",
		10: "Longbottom",
		11: "Potter",
		12: "Snape",
		13: "Adams",
		14: "Willson",
		15: "Daniels",
	}
	rand.Seed(time.Now().UnixNano())
	name := usernames[rand.Intn(16)]
	surname := surnames[rand.Intn(16)]
	birthday := randate().Format("2006-01-02")
	login := surname[0:rand.Intn(len(surname))] + name[0:rand.Intn(len(name))]
	if len(login) < 3 {
		login = login + surname[0:rand.Intn(len(surname))] + name[0:rand.Intn(len(name))]
	}
	rand.Seed(time.Now().Unix())
	minSpecialChar := 2
	minNum := 1
	minUpperCase := 1
	passwordLength := 20
	password := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)

	user := User{
		login,
		name,
		surname,
		birthday,
		password,
	}
	outUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	sendUser(outUser)
}

func randate() time.Time {
	min := time.Date(1940, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2018, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func sendUser(user []byte) {
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
	err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         user,
	})
	if err != nil {
		log.Fatalf("Error publishing message : %s", err)
	}
	log.Printf("User:%s", user)

}
func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
