FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o crawl-service .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/crawl-service .

EXPOSE 8080

CMD ["./crawl-service"]
