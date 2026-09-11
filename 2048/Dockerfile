# Build the Go Binary
FROM --platform=$BUILDPLATFORM golang:alpine AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Podman injects TARGETOS and TARGETARCH dynamically during manifest builds
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /2048 .

# Minimal Runtime
FROM scratch
WORKDIR /app
COPY --from=go-builder /2048 /app/2048
EXPOSE 8080
CMD ["/app/2048"]
