FROM golang:1.13-alpine as builder

RUN apk --update --no-cache add make git g++ linux-headers
# DEBUG
RUN apk add busybox-extras

# Get and build tracer
ADD . /go/src/github.com/vulcanize/tracing-api
WORKDIR /go/src/github.com/vulcanize/tracing-api
RUN make linux

# Build migration tool
WORKDIR /
RUN go get -u -d github.com/pressly/goose/cmd/goose
WORKDIR /go/src/github.com/pressly/goose/cmd/goose
RUN GCO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -tags='no_mysql no_sqlite' -o goose .

# app container
FROM alpine

RUN apk --update --no-cache add postgresql-client

WORKDIR /app

# keep binaries immutable
COPY --from=builder /go/src/github.com/vulcanize/tracing-api/build/tracer-linux /usr/local/bin/tracer
COPY --from=builder /go/src/github.com/pressly/goose/cmd/goose/goose /usr/local/bin/goose
COPY --from=builder /go/src/github.com/vulcanize/tracing-api/startup_script.sh .
COPY --from=builder /go/src/github.com/vulcanize/tracing-api/db/migrations migrations
COPY --from=builder /go/src/github.com/vulcanize/tracing-api/test/enviroment/local/tracing-api/chain.json .

EXPOSE 8080

ENTRYPOINT ["/app/startup_script.sh", "serve"]