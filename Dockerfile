FROM golang:1.23.5-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /app/pod-webhook-mutator ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/pod-webhook-mutator .
COPY certs /root/certs
EXPOSE 443
CMD ["./pod-webhook-mutator"]