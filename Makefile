.PHONY: dev build clean

all: build

run: build
	./notes

build: clean
	go get -d ./...
	go build .

test:
	go test ./...

clean:
	rm -rf notes
