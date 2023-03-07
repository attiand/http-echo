# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN go build -o http-echo

FROM alpine:latest

COPY --from=builder /app/http-echo /http-echo

EXPOSE 8080

CMD ["/http-echo"]
