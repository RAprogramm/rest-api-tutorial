run:
	go run ./cmd/main/app.go

build:
	go build -v ./cmd/main/app.go

.DEFAULT_GOAL := run
