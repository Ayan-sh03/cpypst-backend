# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	air

#create migration
create-migration:
	migrate create -ext sql -dir ./internal/database/migrations $(name)


#migrate 
migrate-up:
	migrate -path ./internal/database/migrations -database "postgres://Ayan:password1234@localhost:5432/cpypst?sslmode=disable" -verbose up
migrate-down:
	migrate -path ./internal/database/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down
.PHONY: all build run test clean

