FROM golang:1.19-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go .

RUN go build -o ./out/logger .

FROM alpine

COPY --from=build_base /tmp/app/out/logger /app/logger

EXPOSE 9095

CMD ["/app/logger"]
