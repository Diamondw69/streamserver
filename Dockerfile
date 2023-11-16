FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o test ./cmd/api

EXPOSE 8080

CMD ["./test"]