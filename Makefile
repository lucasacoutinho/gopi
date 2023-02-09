start:
	@go run ./cmd/api/main.go

migrate:
	@go run ./cmd/database/migrate/main.go migrate

rollback:
	@go run ./cmd/database/migrate/main.go rollback

drop:
	@go run ./cmd/database/migrate/main.go drop

fresh:
	@go run ./cmd/database/migrate/main.go fresh

force:
	@go run ./cmd/database/migrate/main.go force
