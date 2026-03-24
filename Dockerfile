FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY <template> .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/server cmd/server/main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /bin/server /app/server

COPY config/config.yaml /app/config/config.yaml

RUN adduser -D -g '' appuser
USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/server"]
CMD ["-config", "/app/config/config.yaml"]
