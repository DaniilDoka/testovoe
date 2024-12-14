package main

import (
	"testovoe/config"
	"testovoe/internal/server"
	"testovoe/pkg/logger"
	"testovoe/pkg/pg"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	logger := logger.NewLogger()
	server := server.NewServer(logger, cfg)
	pgConn, err := pg.Open(&cfg.Db)
	if err != nil {
		panic(err)
	}
	server.MapRoutes(pgConn)
	server.Run()
}
