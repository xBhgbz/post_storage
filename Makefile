POSTGRES_SETUP=user=postgres password=1234 dbname=postgres host=localhost port=5433 sslmode=disable
MIGRATION_FOLDER=$(CURDIR)/internal/pkg/repository/postgres/migrations

.PHONY: test
test:
	go test -v -tags=integration ./tests/...

.PHONY: app-run
app-run:
	go run ./cmd/storage/main.go

.PHONY: set-up
set-up:
	docker-compose up -d postgres
	timeout 5
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" up
	docker run --name jaegertracing -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest

.PHONY: shut-down
shut-down:
	docker-compose down
	docker stop jaegertracing
	docker remove jaegertracing