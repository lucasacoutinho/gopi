#!/bin/sh
go run ./cmd/database/migrate/main.go create "$@"
