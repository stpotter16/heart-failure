run: build
	@./tmp/main

build:
	@go build -o ./tmp/main cmd/app/main.go

format:
	@go fmt github.com/...
