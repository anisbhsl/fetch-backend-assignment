# Stage 1: Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /fetch-app main.go 

# Stage 2: Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /fetch-app .
EXPOSE 3000
CMD ["./fetch-app"]
