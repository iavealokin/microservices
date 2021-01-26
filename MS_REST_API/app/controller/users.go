package controller

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/iavealokin/MDSite/app/model"

	"github.com/julienschmidt/httprouter"
)

//GetUsers from JSON array to indexDynamic.html
func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем список всех пользователей
	_, err := model.GetUser()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//указываем путь к файлу с шаблоном
	main := filepath.Join("public", "html", "index.html")
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
