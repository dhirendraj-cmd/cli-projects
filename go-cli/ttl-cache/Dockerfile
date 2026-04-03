# stage 1: Build
FROM golang:1.25-alpine AS builder
WORKDIR /app

# copy dependency files
COPY go.mod ./

RUN go mod download

# COPY source code
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ttl-cache ./cmd/main.go


# Runtime
FROM alpine:3.21
WORKDIR /root/
COPY --from=builder /app/ttl-cache .


# start application
ENTRYPOINT [ "./ttl-cache" ]

