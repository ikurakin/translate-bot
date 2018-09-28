FROM golang:alpine3.7 as gobuilder

ARG GITURL="git@github.com"
RUN apk --no-cache upgrade && apk --no-cache add musl-dev gcc git make
RUN go get github.com/golang/dep/cmd/dep
COPY ./ /go/src/github.com/ikurakin/tranlate-bot
RUN cd /go/src/github.com/ikurakin/tranlate-bot/ && GITURL=$GITURL make all

FROM alpine

RUN apk --no-cache upgrade && apk --no-cache add ca-certificates

COPY --from=gobuilder /go/src/github.com/ikurakin/tranlate-bot/bin /opt

WORKDIR /opt

CMD ["./translate-bot"]
