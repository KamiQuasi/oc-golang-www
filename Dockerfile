FROM golang:1.8-alpine

USER nobody

RUN mkdir -p /go/src/github.com/KamiQuasi/oc-golang-www
WORKDIR /go/src/github.com/KamiQuasi/oc-golang-www

COPY . /go/src/github.com/KamiQuasi/oc-golang-www
RUN go install

ENTRYPOINT /go/bin/oc-golang-www

EXPOSE 8024