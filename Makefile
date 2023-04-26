run:
	go run ./cmd/main/app.go

build:
	go build -v -o ./build/app ./cmd/main/app.go 

.DEFAULT_GOAL := run
