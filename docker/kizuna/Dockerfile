FROM golang:1.23-alpine3.20 AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux go build -o kizuna ./cmd/kizuna/main.go

FROM alpine:3.20

COPY --from=builder /app/kizuna /app/kizuna

CMD ["/app/kizuna", "-e", "production"]
