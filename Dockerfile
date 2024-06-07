FROM golang:1.22.4-alpine3.20 AS builder
LABEL authors="kalee"

WORKDIR /workspace

COPY . .
RUN go mod tidy
RUN go build -o quantum-fir quantum-five-in-row-backend

FROM alpine:3.20.0

WORKDIR /workspace
COPY --from=builder /workspace/quantum-fir .

EXPOSE 8080

ENTRYPOINT ["/workspace/quantum-fir"]