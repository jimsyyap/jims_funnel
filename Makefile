# Makefile for running Golang and React services

# Backend (Golang) commands
.PHONY: backend-run
backend-run:
	@echo "Starting Golang backend..."
	@cd backend && go run main.go

.PHONY: backend-build
backend-build:
	@echo "Building Golang backend..."
	@cd backend && go build -o bin/server

# Frontend (React) commands
.PHONY: frontend-run
frontend-run:
	@echo "Starting React frontend..."
	@cd frontend && npm start

.PHONY: frontend-build
frontend-build:
	@echo "Building React frontend..."
	@cd frontend && npm run build

# Combined commands
.PHONY: run
run:
	@echo "Starting backend and frontend..."
	@make backend-run & make frontend-run

.PHONY: build
build:
	@make backend-build
	@make frontend-build

# Install dependencies
.PHONY: install
install:
	@echo "Installing backend dependencies..."
	@cd backend && go mod tidy
	@echo "Installing frontend dependencies..."
	@cd frontend && npm install
