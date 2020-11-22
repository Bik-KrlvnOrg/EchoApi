FROM golang:1.13.5-alpine

ADD . /src

WORKDIR /src

COPY . .

RUN go mod download

RUN go mod verify

EXPOSE 5750

CMD ["go","run","main.go"]