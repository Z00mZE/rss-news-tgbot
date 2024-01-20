FROM golang:alpine as Builder
WORKDIR /bot
COPY . .
RUN go mod download \
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o service ./cmd/bot/main.go

#   Compact application container
FROM alpine:latest
COPY --from=builder /build/service /service
ENTRYPOINT ["/service"]
