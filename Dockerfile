FROM golang:1.19-alpine AS builder

ARG ARCH=amd64

ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH
ENV GO_VERSION 1.19
ENV GO111MODULE on
ENV CGO_ENABLED=0

# Build dependencies
WORKDIR /go/src/
COPY . .
RUN apk update && apk add make git
RUN mkdir /go/src/build
RUN go build -o build/amlogger

# Second stage
FROM alpine:latest
RUN mkdir /var/log/amlogger
RUN chmod 667 /var/log/amlogger
COPY --from=builder /go/src/build/amlogger /usr/local/bin/amlogger
CMD ["/usr/local/bin/amlogger"]
