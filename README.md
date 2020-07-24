# Devices
A GRPC service which provides info about Tidepool supported devices

The service also exposes a REST interface that proxies the requests to the GRPC server using [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).

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
### Code generation

If you make any changes to the proto service definition, you can regenerate the client, server stub, server interface and gateway-proxy 
by running:
```
make generate
```

### Examples

##### List pumps using GRPC
```
grpcurl -plaintext localhost:50051 api.Devices/ListPumps

{
  "pumps": [
    {
      "id": "6678c377-928c-49b3-84c1-19e2dafaff8d",
      "displayName": "Omnipod Horizon",
      "manufacturers": [
        "Insulet"
      ],
      "model": "Omnipod Horizon"
    }
  ]
}
```

##### Get pump by id using GRPC 
```
grpcurl -plaintext -d '{"id":"6678c377-928c-49b3-84c1-19e2dafaff8d"}' localhost:50051 api.Devices/GetPumpById

{
  "pump": {
    "id": "6678c377-928c-49b3-84c1-19e2dafaff8d",
    "displayName": "Omnipod Horizon",
    "manufacturers": [
      "Insulet"
    ],
    "model": "Omnipod Horizon"
  }
}
```

