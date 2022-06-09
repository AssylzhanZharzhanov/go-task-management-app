FROM golang:1.17-alpine as base
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

FROM base AS builder

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./.bin/app ./cmd/api/main.go

FROM alpine:3

COPY --from=builder ./app/.bin/app .

COPY --from=builder ./app/configs ./configs