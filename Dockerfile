FROM golang:1.20.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/app ./cmd/web

FROM alpine:latest


WORKDIR /app

COPY --from=builder /app/bin/app ./bin/app

ENTRYPOINT ["./bin/app"]
