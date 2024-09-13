FROM golang:1.23.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o task-scheduler .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/task-scheduler /app/
EXPOSE 8080
CMD ["./task-scheduler"]
