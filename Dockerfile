FROM golang:1.19 AS base

FROM base AS dev

WORKDIR /opt/app/tibiago

COPY . .

RUN go install github.com/cosmtrek/air@latest && go mod download
# RUN go build .

CMD ["air", "bot"]