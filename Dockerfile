# Use the official Golang image
FROM golang:1.19-alpine as builder

# Set the current working directory inside the container
WORKDIR /usr/src/app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o weather .

FROM alpine:latest

COPY --from=builder /usr/src/app/weather /usr/src/app/weather

EXPOSE 8080

CMD ["/usr/src/app/weather"]

