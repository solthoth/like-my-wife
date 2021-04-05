FROM alpine:3.13.4 as base
RUN apk add ca-certificates
WORKDIR /app

FROM golang:1.16.3-alpine3.13 AS build
RUN apk add --no-cache git
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v
# Build the Go app
RUN go build -o ./out/like-my-wife .

FROM base 
COPY --from=build /build/out/like-my-wfe /app/like-my-wife

# Run the binary program produced by `go install`
CMD ["like-my-wife"]