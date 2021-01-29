FROM golang:1.13-alpine as builder

RUN apk --update --no-cache add make git g++ linux-headers
# DEBUG
RUN apk add busybox-extras

# Get and build tracer
ADD . /go/src/github.com/vulcanize/tracing-api
WORKDIR /go/src/github.com/vulcanize/tracing-api
RUN make linux

# app container
FROM alpine

# keep binaries immutable
COPY --from=builder /go/src/github.com/vulcanize/tracing-api/build/tracer-linux /usr/local/bin/tracer

EXPOSE 8080

ENTRYPOINT ["tracer"]
