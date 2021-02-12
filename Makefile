TAG_COMMIT := $(shell git rev-list --abbrev-commit --tags --max-count=1)
TAG := $(shell git describe --abbrev=0 --tags ${TAG_COMMIT} 2>/dev/null || true)
COMMIT := $(shell git rev-parse --short HEAD)
DATE := $(shell git log -1 --format=%cd --date=format:"%Y%m%d")
VERSION := $(TAG:v%=%)
ifneq ($(COMMIT), $(TAG_COMMIT))
    VERSION := $(VERSION)-next-$(COMMIT)-$(DATE)
endif
ifeq ($(VERSION), )
    VERSION := $(COMMIT)-$(DATA)
endif
ifneq ($(shell git status --porcelain),)
    VERSION := $(VERSION)-dirty
endif
LDFLAGS := -ldflags "-X 'github.com/vulcanize/tracing-api/cmd.version=$(VERSION)'"

#Database
HOST_NAME = localhost
PORT = 5432
NAME = tracing_api
USER = postgres
CONNECT_STRING=postgresql://$(USER)@$(HOST_NAME):$(PORT)/$(NAME)?sslmode=disable

BIN = $(GOPATH)/bin

## Migration tool
GOOSE = $(BIN)/goose
$(BIN)/goose:
	go get -u -d github.com/pressly/goose/cmd/goose
	go build -tags='no_mysql no_sqlite' -o $(BIN)/goose github.com/pressly/goose/cmd/goose

## Apply all migrations not already run
migrate: $(GOOSE)
	$(GOOSE) -table goose_db_version_trace -dir db/migrations postgres "$(CONNECT_STRING)" up

all: clean test linux darwin windows

clean:
	if [ -d "build" ]; then rm -rf "build"; fi

test:
	go vet ./...
	go fmt ./...
	go test ./...

linux:
	if [ ! -d "build" ]; then mkdir "build"; fi
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o build/tracer-linux

darwin:
	if [ ! -d "build" ]; then mkdir "build"; fi
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o build/tracer-darwin

windows:
	if [ ! -d "build" ]; then mkdir "build"; fi
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o build/tracer-windows

## Build docker image
.PHONY: docker-build
docker-build:
	docker build -t vulcanize/tracing-api -f Dockerfile .