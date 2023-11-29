FROM golang:1.21.4-alpine3.18

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8000

CMD ["go", "run", "cmd/main.go"]