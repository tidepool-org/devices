module github.com/tidepool-org/devices/api

go 1.22

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240116215550-a9fa1716bcac
	google.golang.org/grpc v1.60.1
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20240116215550-a9fa1716bcac // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240116215550-a9fa1716bcac // indirect
)

// Resolve GO-2024-2611
replace google.golang.org/protobuf v1.32.0 => google.golang.org/protobuf v1.33.0
