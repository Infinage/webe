FROM --platform=$BUILDPLATFORM golang:alpine AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /ascii-art .

FROM alpine:latest AS runtime
WORKDIR /app

COPY --from=go-builder /ascii-art /app/ascii-art

EXPOSE 8080
CMD ["/app/ascii-art"]
