FROM golang:1.19 AS base

FROM base AS dev
WORKDIR /opt/app/tibiago

RUN go install github.com/cosmtrek/air@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 2345

ENTRYPOINT ["air"]
# CMD ["air", "bot"]

FROM base AS builder
WORKDIR /opt/app/tibiago

COPY . .
RUN go mod download \
    && go mod verify

RUN go build -o tibiago -a .