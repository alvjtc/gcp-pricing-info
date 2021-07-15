FROM golang:1.16-alpine AS build_base

RUN apk add --no-cache git
WORKDIR /go/src/gcp-pricing-info
COPY . .

RUN go mod tidy
RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 go test -v ./internal/...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/server ./cmd/server/

FROM alpine:latest
RUN apk add ca-certificates

COPY --from=build_base /go/src/gcp-pricing-info/out/server /app/server

ENV PORT 8080
EXPOSE $PORT

ENTRYPOINT ["/app/server"]