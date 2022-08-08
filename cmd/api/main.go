package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/LWich/chat-rest-api/internal/app/config"
	delivery "github.com/LWich/chat-rest-api/internal/app/delivery/http"
	"github.com/LWich/chat-rest-api/internal/app/server"
	"github.com/LWich/chat-rest-api/internal/app/service"
	"github.com/LWich/chat-rest-api/internal/app/store"
	"github.com/LWich/chat-rest-api/pkg/auth"
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
	defer db.Close()

	store := store.New(db)

	tokenManager := auth.NewManager(cfg.Auth.SigninKey)

	service := service.New(store, tokenManager, cfg.Auth.AccessTokenTTL, cfg.Auth.RefreshTokenTTL)

	handler := delivery.New(service, tokenManager)
	handler.Init()

	s := server.New(handler)

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
