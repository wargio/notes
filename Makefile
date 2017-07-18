.PHONY: dev build clean

all: dev

dev: build
	./notes

build: clean
	go get ./...
	go build .

test:
	go test ./...

clean:
	rm -rf notes
