package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/hkam0006/ecom-server/internal/env"
	"github.com/jackc/pgx/v5"
)

func main(){
	ctx := context.Background()
	cfg := config{
		addr: ":1212",
		db: dbConfig{
			dsn: env.GetString("GOOSE_STRING", "host=localhost user=postgres password=password dbname=ecom sslmode=disable"),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)
	logger.Info("connected to database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
	}

	h := api.mount()
	if err := api.run(h); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
