package main

import(
	"log"
	"net/http"
	"go-week1-api/internal/handlers"
)

func main(){
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/time", handlers.TimeHandler)

	log.Println("Server running on: 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}