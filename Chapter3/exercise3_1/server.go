package main

import (
	"exercise/Chapter3/exercise3_1/surface"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	surface.CreateSVG(w)
}
