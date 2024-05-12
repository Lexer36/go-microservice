FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go mod download

COPY src/.env .env

RUN go build -o main ./src/main.go

CMD ["./main"]