
FROM golang:1.16-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/app .

# Start fresh from a smaller image
FROM alpine:3.9

COPY --from=build_base /tmp/app/out/app /app/app

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/app"]
