# Stage one: Build
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download
RUN go mod tidy

RUN go build -o main .

# Stage two: Run
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

# Expose 8080 port (Add web-server later)
# EXPOSE 8080

# Run go main binary
CMD ["./main"]
