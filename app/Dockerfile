FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o app

# final stage

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./app"]
