FROM golang:1.17

ADD . /go/src/invite-token

WORKDIR /go/src/invite-token

RUN go mod tidy
RUN go mod vendor
RUN go build

ENTRYPOINT /go/src/invite-token/invite-token

EXPOSE 3000
