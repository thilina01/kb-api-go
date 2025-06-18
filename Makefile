# Makefile for kb-api-go project
# --------------------------------
# This Makefile defines common build, run, test, Docker, and CI/CD tasks.

# Name of the output binary
APP_NAME = kb-api

# Run the Go app locally
run:
	@echo "🔧 Running the application..."
	go run main.go

# Build the Go binary
build:
	@echo "🔨 Building the binary..."
	go build -o $(APP_NAME) main.go

# Run unit tests
test:
	@echo "🧪 Running tests..."
	go test ./...

# Run linting (requires golangci-lint installed)
lint:
	@echo "🔍 Linting code..."
	golangci-lint run

# Generate Swagger/OpenAPI docs (requires swag installed)
swag:
	@echo "📄 Generating Swagger docs..."
	swag init

# Start Docker services (API + MongoDB)
docker-up:
	@echo "🐳 Starting Docker containers..."
	docker compose up --build

# Stop and remove Docker services and volumes
docker-down:
	@echo "🧹 Stopping and removing Docker containers and volumes..."
	docker compose down -v

# Rebuild only the Go service container
docker-rebuild:
	@echo "♻️ Rebuilding only kb-api container..."
	docker compose build kb-api

# Run the application inside the Dev Container (if using VS Code Dev Containers)
dev-run:
	@echo "🚀 Running inside dev container..."
	go run main.go

# Clean build artifacts
clean:
	@echo "🧼 Cleaning build artifacts..."
	rm -f $(APP_NAME)

# Setup project dependencies (tidy modules, generate docs)
setup:
	@echo "📦 Setting up project..."
	go mod tidy
	make swag

# Full CI/CD pipeline locally: lint + test + build
ci:
	@echo "🔁 Running full CI pipeline (lint + test + build)..."
	make lint || true
	make test
	make build

# Prepare artifacts for release (optional)
release:
	@echo "🚀 Preparing release binary..."
	make clean
	make build
	@echo "✅ Release ready: ./$(APP_NAME)"