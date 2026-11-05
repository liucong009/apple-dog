.PHONY: build run test tidy clean

BINARY := apple-dog
CMD := ./cmd/apple-dog

build:
	go build -o bin/$(BINARY) $(CMD)

run: build
	./bin/$(BINARY)

test:
	go test ./...

tidy:
	go mod tidy

clean:
	rm -rf bin/
