FROM golang:1.20

WORKDIR /app

COPY go.mod .

COPY . .
