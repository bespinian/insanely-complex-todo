FROM golang:1.23-alpine AS build

RUN go install github.com/mitranim/gow@latest

WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download

COPY . /app

RUN go build -o tasks main.go

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/tasks .
CMD [ "./tasks" ]
