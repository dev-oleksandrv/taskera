FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ .

RUN go build -o server ./cmd/server

FROM alpine:latest

COPY --from=builder /app/server /server

EXPOSE 8080

CMD ["/server"]
