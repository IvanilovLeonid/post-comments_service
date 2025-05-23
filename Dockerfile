FROM golang:1.23-alpine AS builder

WORKDIR /app
RUN apk --no-cache add bash git make # gcc gettext musl-dev

# dependencies
COPY go.mod go.sum ./
RUN go mod download

# build
COPY . .
RUN go build -o ./bin/app cmd/api/main.go

# run
FROM alpine AS runner

COPY --from=builder /app/bin/app /
COPY docker.env .env
CMD ["/app"]