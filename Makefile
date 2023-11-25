build:
	go build -o ./bin/auth ./cmd/

run: build
	./bin/auth
