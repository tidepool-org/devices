# Devices
A GRPC service which provides info about Tidepool supported devices

The service also exposes a REST interface which proxies the requests to the GRPC server using [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).

## Development

You can start the service locally by running `make start`.

### Prerequisites

Multiple components of this service (grpc server stub, grpc client and grpc gateway proxy) are generated from the protobuf service definition `api/api.proto`. 
This requires the protocol buffer compiler to be installed on the system. On Mac OS you can install it using `brew`:

```
brew install protobuf
```

For other operating systems, you can follow the [official installation instructions](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation).

You can install the grpc and grpc-gateway-proxy generation plugins by running:
```
make install
```
### Generation

If you make any changes to the proto service definition, you can regenerate the client, server stub, server interface and gateway-proxy 
by running:
```
make generate
```
