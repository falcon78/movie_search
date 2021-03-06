FROM golang:1.18.3-alpine3.16

ENV GO111MODULE=on

RUN apk add alpine-sdk bash git --no-cache \
    && git clone https://github.com/vishnubob/wait-for-it.git /tmp/wait-for-it \
    && mv /tmp/wait-for-it/wait-for-it.sh /usr/local/bin/

COPY . /movie

WORKDIR /movie/cmd/server

RUN go get github.com/cosmtrek/air

CMD ["sh", "-c", "go run ."]

EXPOSE 31415
