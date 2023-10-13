package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zhengx55/golang-project-mongodb/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	// New returns a new initialized Router. Path auto-correction, including trailing slashes, is enabled by default.
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
