# STAGE 1: Build your Go application
FROM golang:1.25-alpine AS builder
WORKDIR /app

# Copy and download dependencies
COPY go.mod ./
RUN go mod download

# Copy source and build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ttl-cache ./cmd/main.go

# STAGE 2: Final lightweight image
FROM alpine:3.21
WORKDIR /root/

# Install ttyd directly from Alpine repositories
RUN apk add --no-cache ttyd

# Copy your Go binary from the builder
COPY --from=builder /app/ttl-cache .

# Expose the port (we will set ttyd to use 8080)
EXPOSE 8080

# Run ttyd: 
# -W enables writing (typing)
# -p 8080 sets the port
# ./ttl-cache is your application
# CMD ["ttyd", "-W", "-p", "8080", "./ttl-cache"]

# -W: Allow typing
# -p 8080: Listen on port 8080
# -t titleFixed: Sets a static window title in the browser
# -t theme: A JSON string setting background and foreground colors
# --once: Close the connection after the process exits
# --exit-proxy: Proxies the exit code and prevents the "Press Enter" prompt
# -W: Allow clients to write (type)
# -p 8080: Run on port 8080
# -o: Accept only one client and exit ttyd after the process ends
# -s 9: Sends SIGKILL when exiting to ensure the session terminates cleanly
CMD ["ttyd", "-W", "-p", "8080", "-o", "-s", "9", \
     "-t", "titleFixed=Go TTL Cache Manager", \
     "-t", "theme={'background': '#000000', 'foreground': '#00FF00'}", \
     "./ttl-cache"]



