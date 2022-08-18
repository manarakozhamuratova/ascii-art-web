FROM golang:1.17-alpine AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go

FROM alpine
LABEL authors="Manara & Khamzauly"
COPY --from=builder /build /app
WORKDIR /app
ENTRYPOINT ["/app/main"]
