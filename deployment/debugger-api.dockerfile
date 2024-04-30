FROM golang:1.21.8-alpine AS builder

WORKDIR /build

ADD ../go.mod .

COPY .. .

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]