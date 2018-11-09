package main

import (
	"flag"

	"github.com/tockn/tamabus/server"
)

func main() {
	var port string
	var env string

	flag.StringVar(&port, "port", "8080", "listening port.")
	flag.StringVar(&env, "environment", "development", "environment")

	flag.Parse()

	s := server.NewServer()
	server.Setup(s, "dbconfig.yml", env)
	s.Run(port)

}
