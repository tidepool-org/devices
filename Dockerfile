FROM golang:1.14.4-alpine AS build

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

CMD ["./devices"]
