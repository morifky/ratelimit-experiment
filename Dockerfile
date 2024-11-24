FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main /build/cmd

WORKDIR /dist

RUN cp /build/main .

# Multi-stage docker image
FROM scratch

COPY --from=builder /dist/main /

ENTRYPOINT ["/main"]