start:
	go run cmd/api/main.go | go run cmd/logfmt/main.go

build-migrator:
	@go build -o bin/migrator cmd/database/migrator/main.go
