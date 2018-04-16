package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"api/router"

	"gopkg.in/mgo.v2"
)

type Doc struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	router := router.NewRouter()
	fmt.Printf("_____________________________ GO API _____________________________\n\n")

	fmt.Printf(" ---- RUNTIME INFO ---- \n")
	fmt.Printf("|                      |\n")
	fmt.Printf("| CPU Cores:  %d        |\n", runtime.NumCPU())
	fmt.Printf("| GOMAXPROCS: %d        |\n", runtime.GOMAXPROCS(0))
	fmt.Printf("| Goroutines: %d        |\n", runtime.NumGoroutine())
	fmt.Printf(" ---------------------- \n\n")

	// -------------------------------- MGO TEST --------------------------------

	mongoURI := "mongodb://renan:renan@ds259865.mlab.com:59865/db-test"

	mongoInfo, err := mgo.ParseURL(mongoURI)
	if err != nil {
		panic(err)
	}

	session, err := mgo.DialWithInfo(mongoInfo)
	if err != nil {
		panic(err)
	}

	collection := session.DB("db-test").C("users")

	doc := Doc{"Renan", "renanroberto1@gmail.com"}
	err = collection.Insert(doc)
	if err != nil {
		panic(err)
	}

	// -------------------------------- MGO TEST --------------------------------

	fmt.Println("Server is running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))
}
