api:
	go run ./cmd/api

help:
	go run ./cmd/api --help

frontend-dev:
	cd frontend && npm run dev

frontend-build:
	cd frontend && npm run build