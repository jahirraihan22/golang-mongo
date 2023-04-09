package main

import (
	"fmt"
	"github.com/jahirraihan22/golang-mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	err := http.ListenAndServe("localhost:8081", r)

	if err != nil {
		fmt.Printf("Failed to connect web server ! %v", err)
	}
}

//func getSession() *mgo.Session {
//	s, err := mgo.Dial("mongodb://localhost:27017/golang-mongo")
//	if err != nil {
//		panic(err)
//	}
//	return s
//}

func getSession() (*mgo.Session, error) {
	// Connect to the MongoDB server
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		return nil, err
	}

	// Optional: Set session mode and other options
	session.SetMode(mgo.Monotonic, true)
	session.SetSafe(&mgo.Safe{})

	// Return the session object
	return session, nil
}
