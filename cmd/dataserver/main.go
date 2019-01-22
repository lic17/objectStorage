package main

import (
	"log"
	"net/http"
	"objectStorage/pkg/heartbeat"
	"objectStorage/pkg/objects"
	"os"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
