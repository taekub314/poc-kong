FROM golang:1.23.4-alpine AS builder

WORKDIR /app

RUN apk update && apk add git

COPY go.mod ./
COPY go.sum ./

RUN go mod download

ADD . .

WORKDIR /app
RUN go build -o /ms-service

FROM alpine:3.21

COPY --from=builder /ms-service .

EXPOSE 8080

CMD [ "/ms-service" ]
