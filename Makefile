GOCMD ?= go
GOBUILD = $(GOCMD) build
GOFMT = $(GOCMD)fmt
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
BINARY_NAME = pipemuta
BINARY_UNIX = $(BINARY_NAME)_unix

default: all

all: test build
build:
	$(GOBUILD) -o ./$(BINARY_NAME) -v -ldflags="-X main.VERSION=$(TAG)"

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -rf dist

run: build
	./$(BINARY_NAME)


release:
	goreleaser release --snapshot --clean

fmt:
	$(GOFMT) -w .

dev:
	go get github.com/githubnemo/CompileDaemon && go install github.com/githubnemo/CompileDaemon
	CompileDaemon -build="$(GOBUILD) -o ./$(BINARY_NAME)" -command="./$(BINARY_NAME) -i ./example/pipemuta.yml -t github -o ./example/.github/workflows/github.yaml" -color="true" -exclude-dir=.git -exclude=".#*" -polling -polling-interval 500