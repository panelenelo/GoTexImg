FROM golang:1.26.2-alpine AS builder
WORKDIR /app
COPY . .
# Disable CGO for a statically linked binary (zero external dependencies)
RUN CGO_ENABLED=0 GOOS=linux go build -o appx ./cmd/web/.

FROM scratch
# Necessary so the container has access to the static and html files
COPY --from=builder /app/ui /ui
COPY --from=builder /app/appx /appx
EXPOSE 8181
ENTRYPOINT ["/appx"]
