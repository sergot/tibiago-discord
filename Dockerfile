FROM golang:1.19 AS base

FROM base AS dev

WORKDIR /opt/app/tibiago

COPY . .

RUN go mod download
# RUN go build .

CMD ["go", "run", "main.go"]