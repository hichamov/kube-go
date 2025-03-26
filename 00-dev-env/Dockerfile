FROM golang:1.20.13
WORKDIR /app
COPY main.go /app
RUN go mod init kubernetes
RUN go mod tidy
CMD go run main.go
