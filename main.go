package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/g-leon/Jokes/actions"
	"github.com/markbates/going/defaults"
)

func main() {
	port := defaults.String(os.Getenv("PORT"), "3000")
	log.Printf("Starting Jokes on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), actions.App()))
}
