package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erabxes/golang-react-todo/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on port:9000")

	log.Fatal(http.ListenAndServe(":9000", r))
}
