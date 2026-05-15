FROM golang:1.26.3 AS builder
WORKDIR /build

COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /build/bin/some-app .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/bin/some-app ./

USER 65532:65532
ENTRYPOINT ["/some-app"]
