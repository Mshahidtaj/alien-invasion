# Stage 1: Build Stage
FROM golang:1.20 AS builder

WORKDIR /app

# Copy the source code to the workspace
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o alien-invasion

# Stage 2: Production Stage
FROM alpine:latest

RUN addgroup -S app && adduser -S -G app app

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/alien-invasion .

RUN chown -R app:app /home/app

USER app

# Run the application
CMD ["./alien-invasion"]