build:
	@go build -o bin/novacart-api cmd/main.go

test:
	@go test -v ./..

run: build
	@./bin/novacart-api