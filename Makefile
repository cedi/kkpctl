.PHONY: all build release test bench trace test_all install install_release clean analyze fmt vet tidy lint cyclo dep tools build_dir

VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`
COMMIT=`git rev-parse HEAD`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-X main.version=${VERSION} \
	-X main.date=${BUILD} \
	-X main.commit=${COMMIT} \
	-X main.builtBy=Makefile

LDFLAGS_BUILD=-ldflags "${LDFLAGS}"
LDFLAGS_RELEASE=-ldflags "-s -w ${LDFLAGS}"

OUTPUT_OBJ=-o build/kkpctl

MAIN_GO=./main.go

all: build

build: build_dir tidy analyze
	go build ${LDFLAGS_BUILD} ${OUTPUT_OBJ} ${MAIN_GO}

release: clean build_dir analyze
	go build ${LDFLAGS_RELEASE} ${OUTPUT_OBJ} ${MAIN_GO}

test: build_dir tidy analyze
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

install:
	cp ./build/kkpctl ${GOPATH}/bin/

clean:
	rm -rf ./build
	go clean -cache

analyze: fmt vet lint cyclo

fmt:
	gofmt -w -s -d .

vet: lint cyclo
	go vet ./...

tidy:
	go mod tidy
	go mod verify

lint:
	golint ./...

cyclo:
	gocyclo -avg -over 15 -ignore "_test|Godeps|vendor/" -total .

tools:
	go get -u github.com/go-delve/delve/cmd/dlv@master
	go get -u honnef.co/go/tools/cmd/staticcheck
	go get -u github.com/cweill/gotests/...
	go get -u golang.org/x/lint/golint
	go get -u github.com/fzipp/gocyclo/cmd/gocyclo
	go get -u github.com/amitbet/gorename
	go get -u golang.org/x/tools/cmd/benchcmp

build_dir:
	mkdir -p ./build/
