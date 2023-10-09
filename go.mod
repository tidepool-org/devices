module github.com/tidepool-org/devices

go 1.14

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-playground/validator/v10 v10.3.0
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.0.0-beta.3
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/tidepool-org/devices/api v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.16.0 // indirect
	google.golang.org/grpc v1.30.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/tidepool-org/devices/api => ./api
