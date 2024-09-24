FROM golang:latest
LABEL authors="bombayv"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./cmd/api/main ./cmd/api/main.go

EXPOSE 8080

CMD ["./cmd/api/main"]