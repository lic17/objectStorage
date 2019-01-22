package main

import (
	"log"
	"net/http"
	"objectStorage/pkg/objects"
	"os"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
