FROM golang:1.14 AS builder

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . /app/
RUN CGO_ENABLED=0 go build -o app

FROM alpine:latest

RUN apk add tzdata

WORKDIR /root

COPY --from=builder /app/app /root/
CMD [ "./app" ]
