package main

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

// mongodb+srv://sona:sonapass@sona.xqsd4ki.mongodb.net/?retryWrites=true&w=majority&appName=Sona

func getSession() {
	s, err := mgo.Dial("mongodb+srv://sona:sonapass@sona.xqsd4ki.mongodb.net/?retryWrites=true&w=majority&appName=Sona")
	if err != nil {
		panic(err)
	}
	return s
}
