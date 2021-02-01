package apiserver

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/iavealokin/microservices/MS_REST_API/app/model"
	"github.com/iavealokin/microservices/MS_REST_API/app/store"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)


var (
errIncorrectEmailOrPassword = errors.New("Incorrect email or password")
errUnathorized = errors.New("Unathorized")
)

type ctxKey int8


type server struct {
	router 		 *mux.Router
	logger 		 *logrus.Logger
	store 		 store.Store
}


func newServer(store store.Store) *server{
s :=&server{
	router: 	  mux.NewRouter(),
	logger: 	  logrus.New(),
	store: 		  store,
}
s.configureRouter()
return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	s.router.ServeHTTP(w,r)
}

func (s *server) configureRouter(){
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/addUser", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/dropUser", s.handleUserDrop()).Methods("POST")
	s.router.HandleFunc("/updateUser", s.handleUserUpdate()).Methods("POST")
	s.router.HandleFunc("/changeDelay", s.handleChangeDelay()).Methods("POST")
	s.router.HandleFunc("/users", s.handleGetUsers()).Methods("GET")
}


func (s *server) handleUsersCreate() http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request){
		u := &model.User{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		_ = json.Unmarshal(body, &u)
			if err := s.store.User().Create(u); err != nil{
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, r, http.StatusCreated, u)
	}
}


func (s *server) handleUserDrop() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		u := &model.User{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		_ = json.Unmarshal(body, &u)

			if err := s.store.User().Drop(u); err != nil{
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, r, http.StatusCreated, u)
	}	
}

func (s *server) handleUserUpdate() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		u := &model.User{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		_ = json.Unmarshal(body, &u)
	
			if err := s.store.User().Update(u); err != nil{
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, r, http.StatusCreated, u)
	}	
}


func (s *server) handleGetUsers() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		
	
			users,err := s.store.User().Get(); 
			if err != nil{
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, r, http.StatusCreated, users)
	}	
}

func (s *server) handleChangeDelay() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var delaymap = make(map[string]int)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
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
		s.respond(w, r, http.StatusCreated, "Delay is changed")
	}
}



func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error){
	s.respond(w, r, code, map[string]string{"error": err.Error()})

}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}){
	w.WriteHeader(code)
	if data != nil{
		json.NewEncoder(w).Encode(data)
	}
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}