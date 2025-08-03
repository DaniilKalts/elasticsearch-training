FROM golang:alpine AS builder
RUN adduser -D -u 1000 gopher
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o elasticsearch-training ./cmd/server/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/elasticsearch-training .
COPY --from=builder /app/config ./config
USER 1000
EXPOSE 8080
ENTRYPOINT ["./elasticsearch-training"]
