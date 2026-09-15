FROM golang:1.25-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o traily-backend ./cmd/main.go

FROM golang:1.25-alpine AS dev

WORKDIR /app

RUN go install github.com/air-verse/air@v1.61.7

COPY go.mod go.sum ./

RUN go mod download

EXPOSE 8080

ENTRYPOINT [ "air", "-c", ".air.toml" ]

FROM alpine:3.24

WORKDIR /app

RUN addgroup -S trailyuser && adduser -S -G trailyuser -H -s /sbin/nologin trailyuser

COPY --from=builder --chown=trailyuser:trailyuser /app/traily-backend /app/traily-backend

USER trailyuser

EXPOSE 8080

ENTRYPOINT [ "./traily-backend" ]
