package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/iavealokin/microservices/MS_REST_API/app/model"
	"github.com/iavealokin/microservices/MS_REST_API/app/store"
	"github.com/sirupsen/logrus"
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
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
//	r.GET("/", controller.StartPage)
	//s.router.HandleFunc("/dropUser", s.handleUserDrop()).Methods("POST")
	//("/dropUser", controller.DropUsers)
/*	r.GET("/users", controller.GetUsers)
	r.POST("/updateUser", controller.UpdateUsers)
	r.POST("/addUser", controller.AddUsers)
	r.POST("/changeDelay", changeDelay)
*/
}


func (s *server) handleUsersCreate() http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request){
		u := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(u); err!=nil{
			s.error(w, r, http.StatusBadRequest, err)
			return 
		}
			if err := s.store.User().Create(u); err != nil{
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			//u.Sanitize()
			s.respond(w, r, http.StatusCreated, u)
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