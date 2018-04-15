package main

import (
	"fmt"
	"log"
	"net/http"

	"api/router"
)

func main() {
	router := router.NewRouter()

	fmt.Println("Server is running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))
}
