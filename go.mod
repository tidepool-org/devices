module github.com/tidepool-org/devices

go 1.21.2

require (
	github.com/go-playground/validator/v10 v10.3.0
	github.com/golang/protobuf v1.5.3
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/pkg/errors v0.9.1
	github.com/tidepool-org/devices/api v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.58.2
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/golang/glog v1.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231009173412-8bfb1ae86b6c // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace github.com/tidepool-org/devices/api => ./api
