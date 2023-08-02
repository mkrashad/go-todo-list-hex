FROM golang:1.20.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Install latest version of migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .

RUN go build -o golang-todo-user cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /app /
COPY --from=builder /go/bin/migrate /
 
EXPOSE 8080
CMD ["./app-golang-user"]