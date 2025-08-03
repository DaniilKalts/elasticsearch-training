FROM golang:alpine AS builder
RUN adduser -D -u 1000 gopher
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o elasticsearch-training ./cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o load-to-elastic ./cmd/load-to-elastic/main.go

FROM alpine
WORKDIR /app
RUN apk add --no-cache bash curl
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/elasticsearch-training .
COPY --from=builder /app/load-to-elastic .
USER gopher
EXPOSE 8080
ENTRYPOINT ["./elasticsearch-training"]
