FROM golang:1.23-alpine AS build

WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download

COPY . /app

RUN go build -o tasks main.go
RUN ls

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/tasks .
CMD [ "./tasks" ]
