package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/lucasacoutinho/gopi/internal/api"
	"github.com/lucasacoutinho/gopi/internal/config"
	"github.com/lucasacoutinho/gopi/internal/http/chi"
	"github.com/lucasacoutinho/gopi/user"
	"github.com/lucasacoutinho/gopi/user/db"
)

func main() {
	c := config.Load()

	conn, err := sql.Open(c.DatabaseDriver, c.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	uService := user.NewService(db.New(conn))

	h := chi.Handlers(uService)

	err = api.Start(c, h)
	if err != nil {
		log.Fatalf("error running api: %s", err)
	}
}
