FROM golang:1.20.11 as builder

WORKDIR /go/src/github.com/recall704/memory-limit-scheduler-plugin

COPY . .
RUN make build


FROM debian:stretch-slim

WORKDIR /app

COPY --from=builder /go/src/github.com/recall704/memory-limit-scheduler-plugin/bin/memory-limit-scheduler-plugin /app/memory-limit-scheduler-plugin

CMD ["/app/memory-limit-scheduler-plugin"]
