# Start with a minimal base image
FROM golang:1.20-alpine AS builder

RUN apk add -U --no-cache ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the Go mod and sum files to the working directory
COPY go ./

# Download the dependencies
RUN go mod download

# Build the Go application
RUN go build -o /app/main main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main /main

# Set the entry point for the container
CMD ["/main"]