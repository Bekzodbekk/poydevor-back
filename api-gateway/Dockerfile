# 1. Build bosqichi
FROM golang:1.23.1-alpine AS builder

WORKDIR /app

# Mod fayllarni ko‘chirib, modlarni yuklab olamiz
COPY go.mod go.sum ./
RUN go mod download

# BARCHA fayllarni loyihadan konteynerga ko‘chiramiz
COPY . .

# Build qilamiz
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# 2. Minimal runtime konteyner
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Barcha fayllarni ko‘chiramiz (shu jumladan `config.yaml`)
COPY --from=builder /app ./

# Port ochamiz
EXPOSE 8084

# Dasturni ishga tushiramiz
CMD ["./main"]
