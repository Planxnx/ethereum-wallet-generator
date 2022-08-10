############################
# STEP 1 build executable binary
############################
FROM golang:1.19.0-alpine AS builder
WORKDIR /src/build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=1
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