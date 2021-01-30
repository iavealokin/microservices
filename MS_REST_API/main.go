package main

import (
	"log"
	"net/http"

	"github.com/iavealokin/Microservices/MS_REST_API/app/controller"
	_ "github.com/lib/pq"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	controller.Init()
	routes(r)

	//прикрепляемся к хосту и свободному порту для приема и обслуживания входящих запросов
	//вторым параметром передается роутер, который будет работать с запросами
	err := http.ListenAndServe("0.0.0.0:8082", r)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {
	//путь к папке со внешними файлами: html, js, css, изображения и т.д.
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	//что следует выполнять при входящих запросах указанного типа и по указанному адресу
	r.GET("/", controller.StartPage)
	r.POST("/dropUser", controller.DropUsers)
	r.GET("/users", controller.GetUsers)
	r.POST("/updateUser", controller.UpdateUsers)
	r.POST("/addUser", controller.AddUsers)
	r.POST("/changeDealy", changeDealy)
}

func changeDealy(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	///Func change dealy im serviec MS_Generation
}
