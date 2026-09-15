from golang:1.25-alpine as builder

workdir /app

env CGO_ENABLED=0 GOOS=linux GOARCH=amd64

copy go.mod go.sum ./

run go mod download

copy . .

run go build -o traily ./cmd/main.go

from alpine:3.24

run addgroup -S trailyuser && adduser -S -G trailyuser -H -s /sbin/nologin trailyuser

copy --from=builder --chown=trailyuser:trailyuser /app/traily /app/traily

user trailyuser

entrypoint [ "./traily" ]
