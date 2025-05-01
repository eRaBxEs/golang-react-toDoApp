package main

import (
	"fmt"
	"golang-react-todoapp/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on port:9000")

	log.Fatal(http.ListenAndServe(":9000", r))
}
