FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/huyhvq/lingio
COPY . /usr/local/go/src/github.com/huyhvq/lingio
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/lingio .


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/huyhvq/lingio/build/lingio /bin/lingio
CMD ["lingio", "up"]
