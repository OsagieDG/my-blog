FROM golang:1.20.3-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o bin/app ./cmd/web

CMD ["./bin/app"]
