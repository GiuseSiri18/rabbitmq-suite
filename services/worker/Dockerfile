FROM golang:1.21-alpine

# Install git
RUN apk add --no-cache git

WORKDIR /app

# 1. Initialize the module directly in the container
RUN go mod init worker

# 2. Download the library and ensure it's in the go.mod
RUN go get github.com/streadway/amqp@v1.1.0

# 3. NOW copy the source code (main.go)
COPY main.go ./

# 4. Final tidy to sync everything
RUN go mod tidy

# 5. Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

CMD ["./main"]