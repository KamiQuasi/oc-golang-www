FROM golang:1.8-alpine

USER nobody

RUN mkdir -p /go/src/github.com/KamiQuasi/oc-golang-www
WORKDIR /go/src/github.com/KamiQuasi/oc-golang-www

COPY . /go/src/github.com/KamiQuasi/oc-golang-www
RUN go-wrapper download && go-wrapper install

CMD ["go-wrapper", "run"]
