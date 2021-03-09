.PHONY: all

all: build

build: fmt
	go build -o build/kkpctl ./main.go

test:
	go test ./...

install: build
	cp ./build/kkpctl ${GOPATH}/bin

clean:
	rm -rf ./build

fmt:
	go fmt ./...
