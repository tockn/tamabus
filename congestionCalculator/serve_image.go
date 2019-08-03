package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	if err := http.ListenAndServe(":3000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		if strings.HasPrefix(r.URL.Path, "/raw") {
			http.ServeFile(w, r, "./images/raw/1.png")
		} else if strings.HasPrefix(r.URL.Path, "/result") {
			http.ServeFile(w, r, "./images/result/result.png")
		} else {
			http.NotFound(w, r)
		}
	})); err != nil {
		log.Fatal(err)
	}
}
