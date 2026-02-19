FROM golang:1.24-alpine AS goose-builder

# Install dependencies for goose
RUN apk add --no-cache git
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# -----

FROM golang:1.24-alpine

# Copy goose from builder
COPY --from=goose-builder /go/bin/goose /go/bin/goose

WORKDIR /app
COPY . .

# Run go mod download with the git config in place
RUN go mod download
