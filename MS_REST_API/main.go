package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/iavealokin/microservices/MS_REST_API/app/apiserver"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

var configPath string

func init(){
	flag.StringVar(&configPath,"config-path","configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()
	config:=apiserver.NewConfig()
	_,err := toml.DecodeFile(configPath,config)
	if err!= nil{
		log.Fatal(err)
	}
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
	/*r := httprouter.New()
	controller.Init()
	routes(r)

	//прикрепляемся к хосту и свободному порту для приема и обслуживания входящих запросов
	//вторым параметром передается роутер, который будет работать с запросами
	err := http.ListenAndServe("0.0.0.0:8082", r)
	if err != nil {
		log.Fatal(err)
	}*/
}
/*
func routes(r *httprouter.Router) {
	//путь к папке со внешними файлами: html, js, css, изображения и т.д.
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	//что следует выполнять при входящих запросах указанного типа и по указанному адресу
	r.GET("/", controller.StartPage)
	r.POST("/dropUser", controller.DropUsers)
	r.GET("/users", controller.GetUsers)
	r.POST("/updateUser", controller.UpdateUsers)
	r.POST("/addUser", controller.AddUsers)
	r.POST("/changeDelay", changeDelay)
}
*/
func changeDelay(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var delaymap = make(map[string]int)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.Unmarshal(body, &delaymap)

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
	delay, err := json.Marshal(delaymap)
	if err != nil {
		panic(err)
	}
	err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         delay,
	})
	if err != nil {
		log.Fatalf("Error publishing message : %s", err)
	}
	log.Printf("delay:%s", delay)
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
