FROM golang:1.20.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .
# RUN go build -o api-gw-golang-todo cmd/main.go
RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o app-golang-user cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /app /
COPY --from=builder /go/bin/dlv /

EXPOSE 8080 2345
# CMD ["./api-gw-golang-todo"]
CMD ./dlv --listen=:2345 --headless=true --api-version=2 --log exec ./api-gw-golang-todo