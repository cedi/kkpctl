.PHONY: all

all: build

build: build_dir fmt tidy
	go build -race -o build/kkpctl ./main.go

release: build_dir
	go build -ldflags "-s -w" -o build/kkpctl ./main.go

test: build_dir
	go test -covermode=count -coverprofile=./build/profile.out ./...
	if [ -f ./build/profile.out ]; then go tool cover -func=./build/profile.out; fi

bench: build_dir
	go test -o=./build/bench.test -bench=. -benchmem .
	go test -o=./build/bench.test -cpuprofile=./build/cpuprofile.out .
	go test -o=./build/bench.test -memprofile=./build/memprofile.out .
	go test -o=./build/bench.test -blockprofile=./build/blockprofile.out .
	go test -o=./build/bench.test -mutexprofile=./build/mutexprofile.out ./.

trace: build_dir
	go test -trace=./build/trace.out .
	if [ -f ./build/trace.out ]; then go tool trace ./build/trace.out; fi

test_all: test
	go test all

install: build
	cp ./build/kkpctl ${GOPATH}/bin

install_release: release
	cp ./build/kkpctl ${GOPATH}/bin

clean:
	rm -rf ./build
	go clean -cache

fmt:
	gofmt -w -s -d .

vet: lint
	go vet ./...

tidy:
	go mod tidy
	go mod verify

lint:
	golint ./...

build_dir:
	mkdir -p ./build/