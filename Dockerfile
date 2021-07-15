FROM golang:1.16-alpine AS build_base

RUN apk add --no-cache git
WORKDIR /go/src/gcp-pricing-info
COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o ./out/server ./cmd/server/

FROM alpine:latest
RUN apk add ca-certificates

COPY --from=build_base /go/src/gcp-pricing-info/out/server /app/server

ENV PORT 8080
EXPOSE $PORT

ENTRYPOINT ["/app/server"]