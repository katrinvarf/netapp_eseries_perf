FROM golang:alpine as builder

RUN apk --no-cache add git
RUN go get -u github.com/katrinvarf/san_perf

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /

COPY --from=builder /go/bin/san_perf ./usr/bin/san_perf
CMD ["san_perf", "-config", "/etc/config.yml"]

