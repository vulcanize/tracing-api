FROM golang:1.16-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

WORKDIR /ethereum
RUN git clone https://github.com/vulcanize/go-ethereum.git -b v1.9.25-statediff /ethereum
RUN make geth

# app container
FROM alpine

RUN apk add --no-cache bash

WORKDIR /app

COPY --from=builder /ethereum/build/bin/geth /usr/local/bin/geth
COPY datadir /app/datadir
COPY genesis.json /app
COPY startgeth .
RUN touch /app/pwd && geth --datadir "/app/datadir" init /app/genesis.json

EXPOSE 8545 8546

ENTRYPOINT ["/app/startgeth"]