package main

import (
	"context"
	"flag"
	logger "log"
	"os"

	"github.com/muratmirgun/api/cmd/server"
	"github.com/muratmirgun/api/internal/config"
	"github.com/muratmirgun/api/pkg/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Version set current code version
var Version = "1.1.0"

var configPath = flag.String("config", "./config/local.yaml", "path to the config file")

func main() {
	if err := run(); err != nil {
		logger.Println("error :", err)
		os.Exit(1)
	}
}

func initialize() (*config.Config, log.Logger, error) {
	logger := log.New().With(context.Background(), "version", Version)
	flag.Parse()

	cfg, err := config.Load(*configPath, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		return nil, nil, err
	}

	return cfg, logger, nil
}

func run() error {
	// initialize config and logging
	cfg, logger, err := initialize()
	if err != nil {
		logger.Errorf("failed to initialize: %s", err)
		return err
	}
	// connect to the database
	db, err := sqlx.Connect("mysql", cfg.DB.Dsn)
	if err != nil {
		logger.Errorf("failed to connect to database: %s", err)
		return err
	}

	r := server.Routing(db, logger)
	return server.Start(cfg, r, logger)
}
