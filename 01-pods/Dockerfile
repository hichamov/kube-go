FROM golang:1.21.3
WORKDIR /app
COPY main.go /app
RUN go mod init kubernetes
RUN go mod tidy
CMD go run /app/main.go
