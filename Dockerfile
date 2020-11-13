FROM golang:1.13.5-alpine AS build

WORKDIR /src

COPY . .

RUN go mod download

RUN go mod verify

EXPOSE 5750

RUN go run main.go