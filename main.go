package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"api/router"
)

func main() {
	router := router.NewRouter()
	fmt.Printf("_____________________________ GO API _____________________________\n\n")

	fmt.Printf(" ---- RUNTIME INFO ---- \n")
	fmt.Printf("|                      |\n")
	fmt.Printf("| CPU Cores:  %d        |\n", runtime.NumCPU())
	fmt.Printf("| GOMAXPROCS: %d        |\n", runtime.GOMAXPROCS(0))
	fmt.Printf("| Goroutines: %d        |\n", runtime.NumGoroutine())
	fmt.Printf(" ---------------------- \n\n")

	fmt.Println("Server is running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))
}
