NAME := 2s
REVISION := $(shell git describe --always)
LDFLAGS := -X 'main.Revision=${REVISION}'

clean:
	rm ${GOPATH}/bin/${NAME}
	rm bin/${NAME}

deps:
	go get -d -t -v .

build: deps
	go build -ldflags "${LDFLAGS}" -o bin/${NAME}

test-all: test test-race vet lint

test:
	go test -v -timeout=30s -parallel=4 .

test-race:
	go test -race .

vet:
	go vet *.go

lint:
	@go get github.com/golang/lint/golint
	golint ./...

	
