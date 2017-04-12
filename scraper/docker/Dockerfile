FROM golang:1.8.1-alpine

#/go/src/app should already exist from base image
WORKDIR /go/src/app

COPY . /go/src/app
RUN apk add --no-cache git && go-wrapper download &&  go-wrapper install

CMD ["go-wrapper", "run"]
