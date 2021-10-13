# syntax=docker/dockerfile:1

From golang:1.17.2

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

ADD ./ ./

RUN go build -o /comments-service

CMD ["/comments-service"]