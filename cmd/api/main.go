package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/LWich/chat-rest-api/internal/app/config"
	"github.com/LWich/chat-rest-api/internal/app/handler"
	"github.com/LWich/chat-rest-api/internal/app/server"
	"github.com/LWich/chat-rest-api/internal/app/store"
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

	db, err := newDB(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	store := store.New(db)

	h := handler.NewHelloHandler(store)
	s := server.New(h)

	s.Run(&cfg.Server)
}

func newDB(cfg config.PostgresConfig) (*sql.DB, error) {
	databaseUrl := fmt.Sprintf(
		"postgres://%s/%s?sslmode=%s",
		cfg.PostgresHost,
		cfg.PostgresDbName,
		cfg.PostgresSslMode,
	)
	db, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
