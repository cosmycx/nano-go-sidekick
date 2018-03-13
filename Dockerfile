FROM golang:latest

MAINTAINER _

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go build -o main .

EXPOSE 9077

CMD ["/app/main"]
