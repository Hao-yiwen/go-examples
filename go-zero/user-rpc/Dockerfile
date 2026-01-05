# Build stage
FROM golang:latest AS builder

# Install build dependencies
RUN apk add --no-cache git || apt-get update && apt-get install -y git

WORKDIR /app

# Copy go workspace files first (go.work.sum is optional)
COPY go.work ./
COPY go.work.sum* ./

# Copy all module directories
COPY common/ common/
COPY user-api/ user-api/
COPY user-rpc/ user-rpc/
COPY auth-rpc/ auth-rpc/
COPY role-rpc/ role-rpc/

# Build the user-rpc service
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /user-rpc ./user-rpc/user.go

# Runtime stage
FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy binary from builder
COPY --from=builder /user-rpc .

# Create config directory
RUN mkdir -p /app/etc

EXPOSE 9001

CMD ["./user-rpc", "-f", "/app/etc/user.yaml"]
