FROM golang:alpine as builder

RUN apk --no-cache add git
RUN go install github.com/katrinvarf/netapp_eseries_perf@23880a8

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /

COPY --from=builder /go/bin/netapp_eseries_perf ./usr/bin/netapp_eseries_perf
CMD ["netapp_eseries_perf", "-config", "/etc/config.yml"]

