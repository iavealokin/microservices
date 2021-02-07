package apiserver

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/iavealokin/microservices/MS_REST_API/app/store"
	"github.com/sirupsen/logrus"
)




type webserver struct {
	router 		 *mux.Router
	logger 		 *logrus.Logger
	store 		 store.Store
}


func newWebServer(store store.Store) *webserver{
ws :=&webserver{
	router: 	  mux.NewRouter(),
	logger: 	  logrus.New(),
	store: 		  store,
}
ws.router.StrictSlash(true)
staticDir := "/public/"
ws.router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
ws.configureWebRouter()
return ws
}

func (ws *webserver) ServeHTTP(w http.ResponseWriter, r *http.Request){
	ws.router.ServeHTTP(w,r)
}

func (ws *webserver) configureWebRouter(){
	ws.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	staticDir := "/public/"
	http.FileServer(http.Dir(staticDir))
	ws.router.HandleFunc("/",ws.handleStartPage)
	ws.router.HandleFunc("/login", ws.handleUserLogin)
}


func (ws*webserver) handleStartPage(w http.ResponseWriter, r *http.Request){
	//указываем путь к нужному файлу
	path := filepath.Join("public", "html", "login.html")
	//создаем html-шаблон
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	//выводим шаблон клиенту в браузер
	err = tmpl.ExecuteTemplate(w,"error","")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (ws*webserver) handleUserLogin(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		login:= r.FormValue("login")
		password :=r.FormValue("password")
		err :=ws.store.User().UserLogin(login,password)
		if err!=nil{
			path := filepath.Join("public", "html", "login.html")
			//создаем html-шаблон
			tmpl, err := template.ParseFiles(path)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
		
			//выводим шаблон клиенту в браузер
			err = tmpl.ExecuteTemplate(w,"error","Incorrect login or password!")
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			
		}else{
			users,err := ws.store.User().Get(); 
	if err != nil{
		ws.error(w, r, http.StatusUnprocessableEntity, err)
		return
	}
	//указываем путь к нужному файлу
	path := filepath.Join("public", "html", "home.html")
	//создаем html-шаблон
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	//выводим шаблон клиенту в браузер
	err = tmpl.ExecuteTemplate(w,"users",users)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
		}
			
}


func (ws *webserver) error(w http.ResponseWriter, r *http.Request, code int, err error){
	ws.respond(w, r, code, map[string]string{"error": err.Error()})

}

func (ws *webserver) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}){
	w.WriteHeader(code)
	if data != nil{
		json.NewEncoder(w).Encode(data)
	}
}
