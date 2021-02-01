package controller

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/iavealokin/Microservices/MS_REST_API/app/model"

	"github.com/julienschmidt/httprouter"
)

var user model.User

//GetUsers from JSON array to indexDynamic.html
func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем список всех пользователей
	_, err := model.GetUsersList()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//указываем путь к файлу с шаблоном
	main := filepath.Join("public", "html", "home.html")
	common := filepath.Join("public", "html", "template.html")

	//создаем html-шаблон
	tmpl, err := template.ParseFiles(main, common)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//исполняем именованный шаблон "users", передавая туда массив со списком пользователей
	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

//Init Db fucntion
func Init() {
	err := model.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}
}

//DropUsers API.
func DropUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.Unmarshal(body, &user)
	_ = user.ID
	_, err = model.DropUser(user.ID)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	rw.Write([]byte("User id deleted"))
}

//AddUsers API.
func AddUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.Unmarshal(body, &user)
	_ = user.ID
	_, err = model.AddUser(&user)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	rw.Write([]byte("User id added"))
}

//UpdateUsers API.
func UpdateUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.Unmarshal(body, &user)
	_ = user.ID
	_, err = model.UpdateUser(&user)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	rw.Write([]byte("User id updated"))

}
