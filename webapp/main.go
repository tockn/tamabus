package main

import (
	"flag"
	"log"

	"github.com/tockn/tamabus/webapp/server"
)

func main() {
	var port string
	var env string

	flag.StringVar(&port, "port", "8080", "listening port.")
	flag.StringVar(&env, "environment", "development", "environment")

	flag.Parse()

	s := server.NewServer()
	if err := s.Setup("dbconfig.yml", env); err != nil {
		log.Fatalf("server setup error. %s", err)
	}
	if err := s.Run(port); err != nil {
		log.Fatal(err)
	}
}