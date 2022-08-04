package main

import (
	"flag"
	"log"

	"github.com/LWich/chat-rest-api/internal/app/config"
	"github.com/LWich/chat-rest-api/internal/app/handler"
	"github.com/LWich/chat-rest-api/internal/app/server"
)

var (
	configName string
)

func main() {
	flag.StringVar(&configName, "configName", "config", "set configuration name.")
	flag.Parse()

	v, err := config.LoadConfig(configName)
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ParseConfig(v)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHelloHandler()
	s := server.New(h)

	s.Run(&cfg.Server)
}
