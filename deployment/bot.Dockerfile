FROM golang:alpine as Builder
WORKDIR /bot
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o service ./cmd/bot/main.go

#   Compact application container
FROM alpine:latest
COPY --from=Builder /bot/service /service
ENTRYPOINT ["/service"]
