FROM golang:1.22.5-alpine3.20 as builder

WORKDIR /app

COPY ./ ./

RUN go mod download
RUN go build -o kizuna /app/main.go

FROM golang:1.22.5-alpine3.20

WORKDIR /app

COPY --from=builder /app/kizuna /app/kizuna

CMD ["/app/kizuna"]
