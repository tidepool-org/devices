SHELL = /bin/sh

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get ./...
SERVICE=devices
DIST=dist
BINARY=$(DIST)/$(SERVICE)
DOCKER_REPOSITORY=tidepool/$(SERVICE)

TOOLS_BIN = tools/bin
PROTOC_GEN_GO = $(TOOLS_BIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC = $(TOOLS_BIN)/protoc-gen-go-grpc

PATH:=$(shell pwd)/$(TOOLS_BIN):$(PATH)

all: test build
ci:	test docker-login docker-build docker-push-ci
dist:
		mkdir -p dist
.PHONY: generate
generate: $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC)
			protoc \
        		-I ./api \
        		--experimental_allow_proto3_optional \
        		--go_out=./api --go_opt=paths=source_relative \
        		--go-grpc_out=./api --go-grpc_opt=paths=source_relative \
        		api/api.proto

install:
		go get \
			github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
			github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
			github.com/rakyll/statik
		mkdir -p api/google/api
		curl -sL https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto -o api/google/api/annotations.proto
		curl -sL https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto -o api/google/api/http.proto
		curl -sL https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/httpbody.proto -o api/google/api/httpbody.proto
		curl -sL https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/field_behavior.proto -o api/google/api/field_behavior.proto

$(PROTOC_GEN_GO):
		GOBIN=$(shell pwd)/$(TOOLS_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2

$(PROTOC_GEN_GO_GRPC):
		GOBIN=$(shell pwd)/$(TOOLS_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
build:	dist
		$(GOBUILD) -o $(DIST) ./...
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -rf $(DIST)
start:	build
		TIDEPOOL_DEVICES_CONFIG_FILENAME=`pwd`/devices.yaml ./$(BINARY)
deps:
		$(GOGET) ./...
docker-login:
		@echo "$(DOCKER_PASSWORD)" | docker login --username "$(DOCKER_USERNAME)" --password-stdin
docker-build:
		docker build -t $(SERVICE) .
docker-push-ci:
ifdef TRAVIS_BRANCH
ifdef TRAVIS_COMMIT
ifdef TRAVIS_PULL_REQUEST_BRANCH
	docker tag $(SERVICE) $(DOCKER_REPOSITORY):PR-$(subst /,-,$(TRAVIS_BRANCH))-$(TRAVIS_COMMIT)
	docker push $(DOCKER_REPOSITORY):PR-$(subst /,-,$(TRAVIS_BRANCH))-$(TRAVIS_COMMIT)
else
	docker tag $(SERVICE) $(DOCKER_REPOSITORY):$(subst /,-,$(TRAVIS_BRANCH))-$(TRAVIS_COMMIT)
	docker tag $(SERVICE) $(DOCKER_REPOSITORY):$(subst /,-,$(TRAVIS_BRANCH))-latest
	docker push $(DOCKER_REPOSITORY):$(subst /,-,$(TRAVIS_BRANCH))-$(TRAVIS_COMMIT)
	docker push $(DOCKER_REPOSITORY):$(subst /,-,$(TRAVIS_BRANCH))-latest
endif
endif
endif

.PHONY: ci-generate
ci-generate: generate

.PHONY: ci-build
ci-build: build

.PHONY: ci-test
ci-test: test
