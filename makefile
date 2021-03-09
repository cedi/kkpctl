.PHONY: all

all: build

build:
	go build -o build/kkpctl ./main.go

test:
	go test ./...

install: build
	cp ./build/kkpctl ${GOPATH}/bin

clean:
	rm -rf ./build
