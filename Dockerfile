FROM golang:1.22.2-alpine AS build

WORKDIR /go/src/github.com/tidepool-org/devices
COPY . .

RUN apk --no-cache add git make && \
    make build


FROM alpine:latest AS release

RUN adduser -D tidepool
WORKDIR /home/tidepool
USER tidepool

COPY --from=build --chown=tidepool /go/src/github.com/tidepool-org/devices/dist/devices /go/src/github.com/tidepool-org/devices/devices.yaml ./
ENV TIDEPOOL_DEVICES_CONFIG_FILENAME="/home/tidepool/devices.yaml"

ADD --chown=tidepool https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.3.2/grpc_health_probe-linux-amd64 /bin/grpc_health_probe
RUN chmod +x /bin/grpc_health_probe

CMD ["./devices"]
