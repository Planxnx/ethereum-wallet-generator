############################
# STEP 1 build executable binary
############################
FROM golang:1.12.2-alpine AS builder
RUN apk update && apk add build-base && apk add gcc
WORKDIR /src/build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o main .

############################
# STEP 2 build a small image
############################
FROM alpine:latest
COPY --from=builder /src/build/main /
WORKDIR /
ENV GOGC=
ENV GOMEMLIMIT=
ENTRYPOINT ["./main","-compatible"]