FROM golang:1.15-alpine

RUN apk add --no-cache make gcc musl-dev linux-headers git bash

WORKDIR /app
ENV GOPATH=/go
RUN git clone https://github.com/vulcanize/statediff-migrations /app/statediff
RUN git clone https://github.com/vulcanize/go-ethereum.git -b v1.9.25-statediff /app/ethereum
RUN go get -u -d github.com/pressly/goose/cmd/goose
RUN go build -tags='no_mysql no_sqlite' -o /go/bin/goose github.com/pressly/goose/cmd/goose

RUN cd /app/ethereum && make geth
COPY datadir /app/datadir
COPY poanet.json /app
COPY startgeth .
RUN touch /app/pwd && /app/ethereum/build/bin/geth --networkid "4" --datadir "/app/datadir" init /app/poanet.json

EXPOSE 8545 8546

ENTRYPOINT ["/app/startgeth"]