FROM golang:1.22.8-alpine3.18 as builder

RUN apk add --no-cache make ca-certificates gcc musl-dev linux-headers git jq bash

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

WORKDIR /app

RUN go mod download

ARG CONFIG=config.yml

# build wallet-sign-go with the shared go.mod & go.sum files
COPY . /app/wallet-sign-go

WORKDIR /app/wallet-sign-go

RUN make

FROM alpine:3.18

COPY --from=builder /app/wallet-sign-go/wallet-sign-go /usr/local/bin
COPY --from=builder /app/wallet-sign-go/${CONFIG} /etc/wallet-sign-go/

WORKDIR /app

ENTRYPOINT ["wallet-sign-go"]
CMD ["-c", "/etc/wallet-sign-go/config.yml"]