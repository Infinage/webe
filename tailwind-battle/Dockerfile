# Build the Go Binary
FROM --platform=$BUILDPLATFORM golang:alpine AS go-builder
WORKDIR /app
RUN apk add -U --no-cache ca-certificates

ARG TARGETOS
ARG TARGETARCH

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -ldflags="-s -w" -o /server .

# Minimal Runtime
FROM scratch
WORKDIR /app

# Copy the certs from the builder into the scratch container
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-builder /server /app/server

EXPOSE 8080
CMD ["/app/server"]
