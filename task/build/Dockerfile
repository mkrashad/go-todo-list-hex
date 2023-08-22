FROM golang:1.20.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Install latest version of migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .

# RUN go build -o app-golang-task cmd/main.go
RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o app-golang-task cmd/main.go


FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /app /
COPY --from=builder /go/bin/migrate /
COPY --from=builder /go/bin/dlv /

EXPOSE 8080 2345
CMD ["./app-golang-task"]
# CMD ./dlv --listen=:2345 --headless=true --api-version=2 --log exec ./app-golang-task
