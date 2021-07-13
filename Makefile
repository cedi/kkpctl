.PHONY: all

VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`
COMMIT=`git rev-parse HEAD`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS="-X github.com/cedi/kkpctl/cmd.Version=${VERSION} \
	-X github.com/cedi/kkpctl/cmd.Date=${BUILD} \
	-X github.com/cedi/kkpctl/cmd.Commit=${COMMIT}"
LDFLAGS_BUILD=-ldflags ${LDFLAGS}
LDFLAGS_RELEASE=-ldflags "-s -w" ${LDFLAGS}

all: build

build: build_dir fmt tidy vet
	go build ${LDFLAGS_BUILD} -race -o build/kkpctl ./main.go

release: clean build_dir
	go build ${LDFLAGS_BUILD} -o build/kkpctl ./main.go

test: build_dir fmt tidy vet
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
	cp ./build/kkpctl ${GOPATH}/bin/

install_release: release
	cp ./build/kkpctl ${GOPATH}/bin/

clean:
	rm -rf ./build
	go clean -cache

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

build_dir:
	mkdir -p ./build/