FROM golang:latest AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" -o main .

FROM gcr.io/distroless/static-debian12
WORKDIR /app

# bring in your fully-static binary
COPY --from=builder /app/main /app/main

EXPOSE 80

# use ENTRYPOINT and the absolute path
ENTRYPOINT ["/app/main"]
