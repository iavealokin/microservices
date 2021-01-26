package controller

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/iavealokin/MDSite/app/model"

	"github.com/julienschmidt/httprouter"
)

//GetBillUser func ...
func GetBillUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bills, err := model.GetUserBills()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//указываем путь к файлу с шаблоном
	main := filepath.Join("public", "html", "template.html")
	common := filepath.Join("public", "html", "home.html")

	user, err := model.GetUser()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//создаем html-шаблон
	tmpl, err := template.ParseFiles(main, common)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//исполняем именованный шаблон "home"
	err = tmpl.ExecuteTemplate(rw, "home", struct{ Bills, User interface{} }{bills, user})
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

}
