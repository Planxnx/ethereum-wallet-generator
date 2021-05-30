############################
# STEP 1 build executable binary
############################
FROM golang:1.15.1-alpine AS builder
RUN apk update && apk add build-base && apk add gcc
WORKDIR /src/build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .

############################
# STEP 2 build a small image
############################
FROM alpine:3.12.0
COPY --from=builder /src/build/main /src/app/
WORKDIR /src/app
ENTRYPOINT ["./main"]