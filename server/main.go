package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erabxes/golang-react-todo/router"
	"github.com/rs/cors"
)

func main() {
	r := router.Router()

	// Add CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Your React port
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(r)

	fmt.Println("Starting server on port:9000")
	log.Fatal(http.ListenAndServe(":9000", handler))
}
