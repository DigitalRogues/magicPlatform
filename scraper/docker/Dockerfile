FROM golang:1.7.5-alpine

#/go/src/app should already exist from base image
WORKDIR /go/src/app

COPY . /go/src/app
RUN go-wrapper download &&  go-wrapper install

CMD ["go-wrapper", "run"]
