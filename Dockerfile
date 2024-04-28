FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY cmd /app/cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/web/*.go

FROM alpine:3.18.2

WORKDIR /app

COPY --from=build /app/server /app/server

COPY static /app/static

EXPOSE 8080

CMD [ "/app/server" ]
