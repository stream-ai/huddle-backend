# Use the official Go image as the base image
FROM golang:1.22-alpine as builder

WORKDIR /app

# Copy go.mod and go.sum (dependency management)
COPY go.mod .
COPY go.sum .

# Download project dependencies
RUN go mod download

# Copy source code
COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /app/huddle-backend ./app/

FROM alpine:3.12
COPY --from=builder /app/huddle-backend /app/huddle-backend
ENTRYPOINT ["/app/huddle-backend"]
EXPOSE 80

