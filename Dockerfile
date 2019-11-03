FROM golang:1.13.1 AS build

ENV GO111MODULE=on
ENV GOPATH=""

WORKDIR /go/src/github.com/emilyzhang/lotto-alerts

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -o ./bin/lotto-alerts ./cmd

FROM alpine:3.8

COPY --from=build /go/src/github.com/emilyzhang/lotto-alerts/bin .

RUN ls .

CMD ["./lotto-alerts"]
