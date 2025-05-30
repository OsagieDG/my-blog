FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/app ./cmd/web

FROM alpine:latest AS final

WORKDIR /app

COPY --from=builder /app/bin/app ./bin/app

ENTRYPOINT ["./bin/app"]
