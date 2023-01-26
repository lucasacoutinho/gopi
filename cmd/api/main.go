package main

import (
	"log"

	"github.com/lucasacoutinho/gopi/internal/api"
	"github.com/lucasacoutinho/gopi/internal/config"
	"github.com/lucasacoutinho/gopi/internal/http/chi"
)

func main() {
	c := config.Load()
	h := chi.Handlers()

	err := api.Start(c, h)
	if err != nil {
		log.Fatalf("error running api: %s", err)
	}
}
