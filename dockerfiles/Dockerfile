FROM golang:latest AS builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY . .
RUN go build -o article .

FROM alpine:latest
COPY --from=builder /build/article /usr/local/bin
RUN chmod a+x /usr/local/bin/article
COPY migrates/ /migrates/
COPY testing.yml testing.yml

CMD ["article"]
