.PHONY: build run test tidy clean preview

BINARY := apple-dog
CMD := ./cmd/apple-dog

build:
	go build -o bin/$(BINARY) $(CMD)

run: build
	./bin/$(BINARY) -mode full -plugin weather

preview: build
	./bin/$(BINARY) -list
	./bin/$(BINARY) -activate weather,timer -mode full

test:
	go test ./...

tidy:
	go mod tidy

clean:
	rm -rf bin/
